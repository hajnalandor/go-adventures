import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {  Resize, CopyPaste, Submit } from './data';

const BASE_URI = 'http://localhost:8080/send';
@Injectable({
  providedIn: 'root'
})
export class PostService {

  constructor(private http: HttpClient) { }

  resize(obj: Resize) {
    return this.http.post(BASE_URI, JSON.stringify(obj));
  }

  copypaste(obj: CopyPaste) {
    return this.http.post(BASE_URI, JSON.stringify(obj));
  }
  submit(obj: Submit) {
    return this.http.post(BASE_URI, JSON.stringify(obj));
  }
}
