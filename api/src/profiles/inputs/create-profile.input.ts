import { Field, InputType } from "@nestjs/graphql";

@InputType()
export class CreateProfileInput  {
    @Field((type) => String, {nullable: false})
    email: string;
    @Field((type) => String, {nullable: false})
    name: string;
}