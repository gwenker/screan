import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/operator/startWith';
import 'rxjs/add/operator/switchMap';
import 'rxjs/add/operator/mergeMap';
import 'rxjs/add/operator/toArray';
import { Injectable } from '@angular/core';
import { Action } from '@ngrx/store';
import { Effect, Actions } from '@ngrx/effects';
import { Database } from '@ngrx/db';
import { Observable } from 'rxjs/Observable';
import { defer } from 'rxjs/observable/defer';
import { of } from 'rxjs/observable/of';

import * as sprintAction from '../actions/sprint.action';
import { Sprint } from '../../models/sprint';

import { SprintService } from '../../services/sprint.service'


@Injectable()
export class SprintEffects {
    constructor(private actions$: Actions, private sprintService: SprintService) { }

    @Effect()
    getListSprint$: Observable<Action> = this.actions$
        .ofType(sprintAction.ActionTypes.GET_LIST_SPRINT)
        .switchMap(() =>
            this.sprintService.getListSprints()
                .map((sprints: Sprint[]) => new sprintAction.GetListSprintSuccessAction(sprints))
                .catch(error => of(new sprintAction.GetListSprintFailAction(error)))
        );

}