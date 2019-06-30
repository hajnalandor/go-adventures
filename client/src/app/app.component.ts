import { Component } from '@angular/core';
import { PostService } from './post.service';
import { Data, Dimension } from './data';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  host: {
    '(window:resize)': 'onResize($event)'
  }
})
export class AppComponent {
  title = 'client';
  email = '';
  cardNumber = '';
  securityCode = '';
  t0: number;
  t1: number;
  data = {} as Data;
  resizeFrom = {} as Dimension;


  constructor(private postService: PostService) {
    this.data.websiteurl = location.href.toString();
    this.data.copyAndPaste = false;
    this.resizeFrom.height = window.innerHeight;
    this.resizeFrom.width = window.innerWidth;
  }

  copypaste(event) {
    console.log(event.target.value);
    this.data.copyAndPaste = true;
    this.postService.copypaste(this.data).subscribe(() => {
    });
  }

  startCount() {
    if (!this.t0) {
      this.t0 = performance.now();
    }
  }

  onResize(event){
    const resizeTo = {} as Dimension;
    resizeTo.height = event.target.innerHeight;
    resizeTo.width = event.target.innerWidth;
    this.data.resizeFrom = this.resizeFrom;
    this.data.resizeTo = resizeTo;
    console.log(this.data);
    this.postService.resize(this.data).subscribe(()=> {});
  }

  sendData() {
    this.t1 = performance.now();
    this.data.formCompletionTime = this.t1 - this.t0;
    console.log(this.data);
    this.postService.submit(this.data).subscribe(() => {
    });
  }
}
