import { Field, ID, Int, ObjectType } from '@nestjs/graphql';

@ObjectType()
export class ProfileModel {

    constructor(init?: Partial<ProfileModel>) {
        Object.assign(this, init);
    }

    @Field(() => ID)
    id: string;

    @Field({ nullable: false })
    email: string;

    @Field({ nullable: true })
    name: string;

    @Field({ nullable: true })
    registered: Date;
}