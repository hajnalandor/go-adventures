import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Data } from './data';

const BASE_URI = 'http://localhost:8080/';
@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private http: HttpClient) { }

  resize(obj: Data) {
    return this.http.post(BASE_URI + 'resize', JSON.stringify(obj));
  }

  copypaste(obj: Data) {
    return this.http.post(BASE_URI + 'paste', JSON.stringify(obj));
  }
  submit(obj: Data) {
    return this.http.post(BASE_URI + 'submit', JSON.stringify(obj));
  }
}
