import 'rxjs/add/operator/let';
import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';

import * as fromRoot from './store/reducers';
import * as sprintActions from './store/actions/sprint.action';

import { Sprint } from './models/sprint';

@Component({
  selector: 'screan-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent {
  title = 'app works!';
  
  sprints$: Observable<Array<Sprint>>;

  constructor(private store: Store<fromRoot.State>) {
    this.sprints$ = store.select(fromRoot.getSprints);
    this.store.dispatch(new sprintActions.GetListSprintAction());
  }
  
  getRandomInt():Number {
    return Math.floor(Math.random() * 100) + 1;
  }

}
