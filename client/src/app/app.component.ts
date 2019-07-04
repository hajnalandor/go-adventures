import { Component } from '@angular/core';
import { PostService } from './post.service';
import { Dimension, Resize, CopyPaste, Submit } from './data';

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
  resize = {} as Resize;
  copyPaste = {} as CopyPaste;
  submit = {} as Submit;
  resizeFrom = {} as Dimension;


  constructor(private postService: PostService) {
    this.resize.websiteurl = location.href.toString();
    this.copyPaste.websiteurl = location.href.toString();
    this.submit.websiteurl = location.href.toString();

    this.resizeFrom.height = window.innerHeight;
    this.resizeFrom.width = window.innerWidth;
  }

  copypaste(event) {
    console.log(event.target.value);
    this.copyPaste.eventType = 'copyAndPaste';
    this.copyPaste.pasted = true;
    this.copyPaste.formId = event.target.id;
    this.postService.copypaste(this.copyPaste).subscribe(() => {
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
    this.resize.resizeFrom = this.resizeFrom;
    this.resize.resizeTo = resizeTo;
    this.resize.eventType = 'resize';
    console.log(this.resize);
    this.postService.resize(this.resize).subscribe(() => {});
  }

  sendData() {
    this.submit.eventType = 'timeTaken';
    this.t1 = performance.now();
    this.submit.time = Math.round((this.t1 - this.t0) / 1000);
    console.log(this.submit);
    this.postService.submit(this.submit).subscribe(() => {
    });
  }
}
