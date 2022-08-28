import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { MusicalEventPage } from './musical-event.page';

const routes: Routes = [
  {
    path: '',
    component: MusicalEventPage
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class MusicalEventPageRoutingModule {}
