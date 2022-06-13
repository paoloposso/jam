import { Args, Mutation, Query, Resolver } from "@nestjs/graphql";
import { AddInstrumentsInput } from "./inputs/add-instruments.input";
import { CreateProfileInput } from "./inputs/create-profile.input";
import { ProfileModel } from "./models/profile.model";
import { ProfileService } from "./profile.service";

@Resolver(of => ProfileModel)
export class ProfileResolver {
  
  constructor(private service: ProfileService) {}

  @Query(returns => ProfileModel, { name: 'getProfileByEmail' })
  getProfileByEmail(@Args('email', { type: () => String }) email: string) {
    return this.service.getProfileByEmail(email);
  }

  @Query(returns => ProfileModel, { name: 'getProfileById' })
  getProfileById(@Args('id', { type: () => String }) email: string) {
    return this.service.getProfileById(email);
  }

  @Mutation(returns => String)
  async createProfile(@Args({ name: 'input', type: () => CreateProfileInput }) input: CreateProfileInput) {
    return this.service.create(new ProfileModel(input));
  }

  @Mutation(returns => String)
  async addInstruments(@Args({ name: 'input', type: () => AddInstrumentsInput }) input: AddInstrumentsInput) {
    return this.service.addInstruments(input.id, input.instruments);
  }
}