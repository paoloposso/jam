import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { Profile } from '../pages/profile/models/profile';
import { retry, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {

  url = 'http://localhost:8000/users';

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  }

  constructor(private httpClient: HttpClient) { }

  saveBasicData(email: string, name: string, password: string): Observable<Profile> {
    return this.httpClient
      .post<Profile>(this.url, {email, name, password}, this.httpOptions)
      .pipe(retry(2), catchError(this.handleError));
  }

  save(profile: Profile): Observable<Profile> {
    return this.httpClient
      .post<Profile>(this.url, profile, this.httpOptions)
      .pipe(retry(2), catchError(this.handleError));
  }

  handleError(error: HttpErrorResponse) {
    let errorMessage = '';
    if (error.error instanceof ErrorEvent) {
      errorMessage = error.error.message;
    } else {
      // server side error
      errorMessage = `Error code: ${error.status}, message: ${error.message}`;
    }
    console.log(errorMessage);
    return throwError(errorMessage);
  };

}