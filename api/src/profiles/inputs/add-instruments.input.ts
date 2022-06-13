import { Field, InputType } from "@nestjs/graphql";

@InputType()
export class AddInstrumentsInput  {
    @Field((type) => String, {nullable: false})
    id: string;
    @Field((type) => [String], {nullable: false})
    instruments: string[];
}
