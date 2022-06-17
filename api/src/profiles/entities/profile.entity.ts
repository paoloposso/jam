import { Field, ObjectType } from '@nestjs/graphql';
import { LocationEntity } from './location.entity';

@ObjectType()
export class ProfileEntity {

    @Field()
    id: string;

    @Field({ nullable: false })
    email: string;

    @Field({ nullable: true })
    name: string;

    @Field({ nullable: true })
    registered: Date;

    @Field(_type => [String], {nullable: true })
    styles: string[];

    @Field(_type => [String], { nullable: true })
    instruments: string[];

    @Field({ nullable: true })
    location: LocationEntity;
}