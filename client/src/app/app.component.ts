import 'rxjs/add/operator/let';
import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs/Observable';

import * as fromRoot from './store/reducers';
import * as sprintActions from './store/actions/sprint.action';

import { Sprint } from './models/sprint';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app works!';
  
  sprints$: Observable<{ [id: string]: Sprint; }>;

  constructor(private store: Store<fromRoot.State>) {
    this.sprints$ = store.select(fromRoot.getSprints);
    this.store.dispatch(new sprintActions.GetListSprintAction());
  }
}
