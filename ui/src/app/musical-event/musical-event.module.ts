import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';

import { MusicalEventPageRoutingModule } from './musical-event-routing.module';

import { MusicalEventPage } from './musical-event.page';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    MusicalEventPageRoutingModule
  ],
  declarations: [MusicalEventPage]
})
export class MusicalEventPageModule {}
