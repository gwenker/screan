import { Component, OnInit, Input } from '@angular/core';
import { Sprint } from '../../models/sprint';

@Component({
  selector: 'screan-sprint-box',
  templateUrl: './sprint-box.component.html',
  styleUrls: ['./sprint-box.component.css']
})
export class SprintBoxComponent implements OnInit {

  @Input()
  sprint: Sprint;

  constructor() { }

  ngOnInit() {
  }

}
