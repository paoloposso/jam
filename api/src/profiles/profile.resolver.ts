import { Args, Mutation, Query, Resolver } from "@nestjs/graphql";
import { AddInstrumentsInput } from "./inputs/add-instruments.input";
import { CreateProfileInput } from "./inputs/create-profile.input";
import { ProfileEntity } from "./entities/profile.entity";
import { ProfileService } from "./profile.service";

@Resolver(of => ProfileEntity)
export class ProfileResolver {
  
  constructor(private service: ProfileService) {}

  @Query(returns => ProfileEntity, { name: 'getProfileByEmail' })
  getProfileByEmail(@Args('email', { type: () => String }) email: string) {
    return this.service.getProfileByEmail(email);
  }

  @Query(returns => ProfileEntity, { name: 'getProfileById' })
  getProfileById(@Args('id', { type: () => String }) id: string) {
    return this.service.getProfileById(id);
  }

  @Mutation(returns => String)
  async createProfile(@Args({ name: 'input', type: () => CreateProfileInput }) input: CreateProfileInput) {
    return this.service.create(Object.assign(new ProfileEntity(), input));
  }

  @Mutation(returns => String)
  async addInstruments(@Args({ name: 'input', type: () => AddInstrumentsInput }) input: AddInstrumentsInput) {
    return this.service.addInstruments(input.id, input.instruments);
  }
}