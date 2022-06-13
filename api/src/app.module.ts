import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { DirectiveLocation, GraphQLDirective } from 'graphql';
import { ProfileModule } from './profiles/profile.module';
import { MongooseModule } from '@nestjs/mongoose';
import * as dotenv from "dotenv";

dotenv.config();

@Module({
  imports: [
    ProfileModule,
    MongooseModule.forRoot('mongodb://localhost:27017/jam'),
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: 'schema.gql',
      installSubscriptionHandlers: true,
      playground: true,
      buildSchemaOptions: {
        directives: [
          new GraphQLDirective({
            name: 'upper',
            locations: [DirectiveLocation.FIELD_DEFINITION],
          }),
        ],
      },
    }),
  ]
})
export class AppModule {}