import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IonicModule } from '@ionic/angular';
import { ProfilePage } from './profile.page';
import { ProfilePageRoutingModule } from './profile-routing.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ProfileService } from '../../services/profile.service';

@NgModule({
  imports: [
    CommonModule,
    IonicModule,
    ProfilePageRoutingModule,
    FormsModule,
    ReactiveFormsModule
  ],
  declarations: [ProfilePage],
  providers: [ProfileService]
})
export class ProfilePageModule {}
