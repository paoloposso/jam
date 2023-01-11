import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { Profile } from './models/profile';
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
