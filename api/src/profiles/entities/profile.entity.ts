import { Field, ID, Int, ObjectType } from '@nestjs/graphql';
import { LocationEntity } from './location.entity';

@ObjectType()
export class ProfileEntity {

    constructor(init?: Partial<ProfileEntity>) {
        Object.assign(this, init);
    }

    @Field()
    id: string;

    @Field({ nullable: false })
    email: string;

    @Field({ nullable: true })
    name: string;

    @Field({ nullable: true })
    registered: Date;

    @Field({ nullable: true })
    location: LocationEntity;
}