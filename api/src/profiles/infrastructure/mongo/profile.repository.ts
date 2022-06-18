import { Injectable } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model } from "mongoose";
import { LocationEntity } from "../../entities/location.entity";
import { ProfileEntity } from "../../entities/profile.entity";
import { IProfileRepository } from "../../profile.repository";
import { ProfileDocument, ProfileModel, ProfileSchema } from "./profile.schema";

@Injectable()
export class ProfileRepository implements IProfileRepository {

    constructor(
        @InjectModel('Profile')
        private model: Model<ProfileDocument>) {}

    public async getByEmail(email: string): Promise<ProfileEntity> {
        let result = new ProfileEntity();

        let profileData = await this.model.findOne({email});
        
        if (profileData) {
            result.id = profileData._id;
            result.email = profileData.email;
            result.name = profileData.name;
            result.location = profileData.location;
        }
        
        return result;
    }

    public async getById(id: string): Promise<ProfileEntity> {
        let result = new ProfileEntity();

        let profileData = await this.model.findById(id);

        if (profileData) {
            result.id = profileData._id;
            result.email = profileData.email;
            result.name = profileData.name;
            result.location = profileData.location;
        }

        return result;
    }

    public async create(profile: ProfileEntity): Promise<string> {
        return (await new this.model(profile).save())._id;
    }

    public async saveLocation(profileId: string, location: LocationEntity): Promise<string> {
        const profileModel = await this.model.findById(profileId);
        Object.assign(profileModel.location, location);
        let result = await profileModel.save();
        return result._id;
    }

    public async addInstruments(id: string, instruments: string[]) {
        const model = await this.model.findById(id);
        model.instruments = [];
        model.instruments.push(...instruments);
        return (await model.save())._id.toString();
    }

    public async addStyles(id: string, styles: string[]) {
        const model = await this.model.findById(id);
        model.styles = [];
        model.styles.push(...styles);
        return (await model.save())._id.toString();
    }
}