import * as moment from "moment";
import { Model } from "mongoose";
import { ProfileEntity } from "./entities/profile.entity";
import { ProfileDocument } from "./infrastructure/mongo/profile.schema";
import { ProfileRepository } from "./infrastructure/mongo/profile.repository";
import { ProfileService } from "./profile.service";
import { IProfileRepository } from "./profile.repository";

describe('ProfileService', () => {
    let service: ProfileService;
    let repository: IProfileRepository;

    beforeAll(() => {
        repository = new ProfileRepository({} as Model<ProfileDocument, {}, {}>);
        service = new ProfileService(repository);
    });

    describe('getProfileByEmail', () => {
        it ('should return a profile', async () => {
            jest.spyOn(repository, 'getByEmail').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve({email: 'pvictorsys@gmail.com'} as ProfileEntity);
                });
            });

            expect((await service.getProfileByEmail('pvictorsys@gmail.com')).email).toEqual('pvictorsys@gmail.com');
        });
    });

    describe('getProfileById', () => {
        it ('should return a profile', async () => {
            jest.spyOn(repository, 'getById').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve({email: 'pvictorsys@gmail.com'} as ProfileEntity);
                });
            });

            expect((await service.getProfileById('111Bhjsyd6&85')).email).toEqual('pvictorsys@gmail.com');
        });
    });

    describe('createProfile', () => {
        it ('should return id for new profile', async () => {
            jest.spyOn(repository, 'create').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve('id123');
                });
            });

            let profile = new ProfileEntity();
            profile.email = 'test@test.com';
            profile.name = 'test';

            let respose = await service.create(profile);

            expect(respose.length).toBeGreaterThan(1);
        });

        it ('should fail when email is empty', async () => {
            jest.spyOn(repository, 'create').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve('id123');
                });
            });

            let profile = new ProfileEntity();
            profile.email = 'pvictorsys@gmail.com';
            profile.name = '';

            try {

                await service.create(profile);
                fail('should have validated blank name');
            } catch (e) {
                let err = e as Error;
                expect(err.message.toLowerCase().includes('name')).toBeTruthy();
            }
        });

        it ('should fail when email is blank', async () => {
            jest.spyOn(repository, 'create').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve('id123');
                });
            });

            let profile = new ProfileEntity();
            profile.email = '';

            try {
                await service.create(profile);
                fail('should have validated email');
            } catch (e) {
                let message = ((e as Error).message);
                expect(message.toLowerCase().includes('email is required')).toBeTruthy();
            }
        });

        it ('should return id after adding styles', async () => {
            let id = 'abcd1234';

            jest.spyOn(repository, 'addStyles').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve(id);
                });
            });

            let profile = {} as ProfileEntity;
            profile.id = id;
            profile.styles = ['jazz', 'samba'];

            let result = await service.addStyles(id, profile.styles);
            expect(id).toEqual(result);
        });

        it ('should fail when email is blank', async () => {
            let id = 'abcd1234';

            jest.spyOn(repository, 'create').mockImplementation(() => {
                return new Promise((resolve, _reject) => {
                    resolve('id123');
                });
            });

            let styles: string[] = [];

            try {
                await service.addInstruments(id, styles);
                fail('should have validated styles');
            } catch (e) {
                let message = ((e as Error).message);
                expect(message.toLowerCase().includes('is required')).toBeTruthy();
            }
        });
    });
});