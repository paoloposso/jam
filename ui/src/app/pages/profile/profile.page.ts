import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormArray, FormControl } from "@angular/forms";
import { Profile } from './models/profile';
import { ProfileService } from '../../services/profile.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {

  profileForm: FormGroup;
  isSubmitted: boolean = false;

  constructor(
    private formBuilder: FormBuilder, 
    private service: ProfileService) { }

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
    this.isSubmitted = true;
    if (this.profileForm.valid) {
      let profile = new Profile();

      profile.email = this.profileForm.get('email').value;
      profile.name = this.profileForm.get('name').value;
      profile.gender = this.profileForm.get('gender').value;

      this.service.save(profile).subscribe(
        p => console.log(p),
        err => alert(err)
      );
    }
  }

  get errorControl() {
    return this.profileForm.controls;
  }
}
