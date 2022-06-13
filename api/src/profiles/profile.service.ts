import { ProfileModel } from "./models/profile.model";
import { ProfileRepository } from "./mongo/profile.repository";
import { Injectable } from "@nestjs/common";
import * as moment from "moment";

@Injectable()
export class ProfileService {

    constructor(private repository: ProfileRepository) {}
    
    public async getProfileByEmail(email: string): Promise<ProfileModel> {
        return this.repository.getByEmail(email);
    }

    public async getProfileById(id: string): Promise<ProfileModel> {
        return this.repository.getById(id);
    }

    public async create(profile: ProfileModel): Promise<string> {
        let errors: string[] = [];

        this.validateCreateProfileInput(profile, errors);

        profile.registered = this.getUtcDate();

        return this.repository.create(profile);
    }

    public async addInstruments(profileId: string, instruments: string[]) {
        return this.repository.addInstruments(profileId, instruments);
    }

    getUtcDate() : Date {
        return moment.utc().toDate();
    }

    private validateCreateProfileInput(profile: ProfileModel, errors: string[]) {
        if (!profile.email || profile.email.length === 0) {
            errors.push('Email is required');
        }
        if (!profile.name || profile.name.length === 0) {
            errors.push('Name is required');
        }
        if (errors.length > 0) {
            throw new Error(errors.join(';'));
        }
    }
}