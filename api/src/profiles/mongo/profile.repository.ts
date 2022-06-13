import { Injectable } from "@nestjs/common";
import { InjectModel } from "@nestjs/mongoose";
import { Model } from "mongoose";
import { ProfileModel } from "../models/profile.model";
import { ProfileDocument } from "./profile.schema";

@Injectable()
export class ProfileRepository {

    constructor(
        @InjectModel('Profile')
        private model: Model<ProfileDocument>) {}

    public async getByEmail(email: string): Promise<ProfileModel> {
        let result: ProfileModel;
        Object.assign(result, await this.model.findOne({ email }).exec());
        return result;
    }

    public async getById(id: string): Promise<ProfileModel> {
        let result: ProfileModel;
        Object.assign(result, await this.model.findById(id).exec());
        return result;
    }

    public async create(profile: ProfileModel): Promise<string> {
        return (await new this.model(profile).save())._id.toString();
    }

    public async addInstruments(id: string, instruments: string[]) {
        const model = await this.model.findById(id);
        model.instruments = [];
        model.instruments.push(...instruments);
        return (await model.save())._id.toString();
    }
}