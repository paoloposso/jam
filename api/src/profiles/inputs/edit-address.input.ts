import { Field, InputType } from "@nestjs/graphql";

@InputType()
export class EditAddressInput  {
    @Field((type) => String, {nullable: false})
    id: string;

    @Field((type) => [Number], {nullable: false})
    coordinates: number[]

    @Field((type) => String, {nullable: false})
    fullAddress: string
}
