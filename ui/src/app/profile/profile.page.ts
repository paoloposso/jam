import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormArray, FormControl } from "@angular/forms";
import { Profile } from './profile';
import { ProfileService } from './profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {

  profileForm: FormGroup;

  constructor(private formBuilder: FormBuilder, private service: ProfileService) { }

  ngOnInit() {
    this.profileForm = this.formBuilder.group({
      name: new FormControl('', Validators.required),
      email: new FormControl('', Validators.required),
      gender: new FormControl('', Validators.required),
      guitar: new FormControl(false),
      bass: new FormControl(false),
      drums: new FormControl(false),
      keyboard: new FormControl(false),
    });
  }

  onSubmit() {
    let m = new Profile();
    m.name = "Paolo";
    this.profileForm.patchValue(m);
  }
}
