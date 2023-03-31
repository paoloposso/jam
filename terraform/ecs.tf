# Create a VPC
resource "aws_vpc" "jam-api" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "jam-api-vpc"
  }
}

# Create an Internet Gateway
resource "aws_internet_gateway" "jam-api" {
  vpc_id = aws_vpc.jam-api.id

  tags = {
    Name = "jam-api-igw"
  }
}

resource "aws_route_table" "jam-api" {
  vpc_id = aws_vpc.jam-api.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.jam-api.id
  }

  tags = {
    Name = "jam-api-rt"
  }
}

# Create a public subnet
resource "aws_subnet" "public" {
  cidr_block        = "10.0.1.0/24"
  vpc_id            = aws_vpc.jam-api.id
  availability_zone = "us-east-2a"

  tags = {
    Name = "jam-api-public-subnet"
  }
}

# Create a private subnet
resource "aws_subnet" "private" {
  cidr_block        = "10.0.2.0/24"
  vpc_id            = aws_vpc.jam-api.id
  availability_zone = "us-east-2b"

  tags = {
    Name = "jam-api-private-subnet"
  }
}

resource "aws_route_table_association" "jam-api" {
  subnet_id      = aws_subnet.public.id
  route_table_id = aws_route_table.jam-api.id
}

# Create a security group for the ECS tasks
resource "aws_security_group" "ecs_tasks" {
  name_prefix = "jsg-"
  vpc_id      = aws_vpc.jam-api.id
}

resource "aws_security_group" "alb_sg" {
  name_prefix = "jsg-"
  vpc_id      = aws_vpc.jam-api.id
}

resource "aws_lb" "jam-api" {
  name               = "jam-api-lb"
  internal           = false
  load_balancer_type = "application"
  subnets            = [aws_subnet.private.id, aws_subnet.public.id]
  security_groups    = [aws_security_group.alb_sg.id]

  tags = {
    Name = "jam-api-lb"
  }
}

resource "aws_lb_target_group" "jam-api" {

  name_prefix      = "jtg-"
  port             = 80
  protocol         = "HTTP"
  target_type      = "ip"
  vpc_id           = aws_vpc.jam-api.id

  health_check {
    path     = "/"
    interval = 30
    timeout  = 10
  }
}

resource "aws_lb_listener" "jam-api" {
  load_balancer_arn = aws_lb.jam-api.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    target_group_arn = aws_lb_target_group.jam-api.arn
    type             = "forward"
  }
}

resource "aws_lb_listener_rule" "jam-api" {
  listener_arn = aws_lb_listener.jam-api.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.jam-api.arn
  }

  condition {
    path_pattern {
      values = ["/*"]
    }
  }
}


resource "aws_ecs_cluster" "jam_ecs_cluster" {
  name = "jam-ecs-cluster"
}

resource "aws_ecs_service" "jam-api" {
  name            = "jam-api"
  cluster         = aws_ecs_cluster.jam_ecs_cluster.id
  task_definition = aws_ecs_task_definition.jam-api.arn

  launch_type = "FARGATE"

  desired_count   = 1

  network_configuration {
    subnets          = [aws_subnet.private.id, aws_subnet.public.id]
    security_groups  = [aws_security_group.ecs_tasks.id]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.jam-api.arn
    container_name   = "jam-api"
    container_port   = 5500
  }
}


resource "aws_ecs_task_definition" "jam-api" {
  family                   = "jam-api"
  container_definitions    = jsonencode([
    {
      name      = "jam-api"
      image     = "docker.io/pvictosys/jam-api:latest"
      portMappings = [
        {
          containerPort = 5500
          hostPort = 5500
        }
      ]
    }
  ])

  requires_compatibilities = [
    "FARGATE"
  ]

  network_mode = "awsvpc"

  memory        = 512
  cpu           = 256

#   execution_role_arn = aws_iam_role.ecs_task_execution_role.arn
#   task_role_arn      = aws_iam_role.ecs_task_role.arn
}
