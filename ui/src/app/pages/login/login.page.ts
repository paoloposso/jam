import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginService } from '../../services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.page.html',
  styleUrls: ['./login.page.scss'],
})
export class LoginPage implements OnInit {
  email: string;
  password: string;

  constructor(private service: LoginService, private router: Router) {}

  ngOnInit() {}

  loginUser(){
    this.service.login(this.email, this.password);
    this.email = "";
    this.password = "";
  }

  register(){
    this.router.navigate(['register']);
  }
}
