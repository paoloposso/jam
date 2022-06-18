import { Field, InputType } from "@nestjs/graphql";

@InputType()
export class AddStylesInput  {
    @Field((type) => String, {nullable: false})
    id: string;
    @Field((type) => [String], {nullable: false})
    styles: string[];
}
