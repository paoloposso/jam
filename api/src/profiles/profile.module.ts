import { Module } from "@nestjs/common";
import { MongooseModule } from "@nestjs/mongoose";
import { ProfileSchema } from "./mongo/profile.schema";
import { ProfileRepository } from "./mongo/profile.repository";
import { ProfileResolver } from "./profile.resolver";
import { ProfileService } from "./profile.service";

@Module({
    imports: [
        MongooseModule.forFeature([{
            name: 'Profile',
            schema: ProfileSchema
        }])
    ],
    providers: [
        ProfileResolver,
        ProfileService,
        ProfileRepository,
        {
            provide: 'IProfileRepository',
            useClass: ProfileRepository
        }   
    ],
})
export class ProfileModule {}
