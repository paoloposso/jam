import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from "@angular/forms";

@Component({
  selector: 'app-profile',
  templateUrl: './profile.page.html',
  styleUrls: ['./profile.page.scss'],
})
export class ProfilePage implements OnInit {

  ionicForm: FormGroup;

  constructor(public formBuilder: FormBuilder) { }

  ngOnInit() {
  }

  getDate(e) {
    let date = new Date(e.target.value).toISOString().substring(0, 10);
    this.ionicForm.get('dob').setValue(date, {
       onlyself: true
    })
  }

  submitForm() {
    console.log(this.ionicForm.value)
  }
}
