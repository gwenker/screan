import { Action } from '@ngrx/store';
import { Sprint } from '../../models/sprint';
import { type } from '../util';

/**
 * For each action type in an action group, make a simple
 * enum object for all of this group's action types.
 * 
 * The 'type' utility function coerces strings into string
 * literal types and runs a simple check to guarantee all
 * action types in the application are unique. 
 */
export const ActionTypes = {
    GET_LIST_SPRINT: type('[Sprint] get list'),
    GET_LIST_SPRINT_SUCCESS: type('[Sprint] get list success'),
    GET_LIST_SPRINT_FAIL: type('[Sprint] get list fail')
};


/**
 * Every action is comprised of at least a type and an optional
 * payload. Expressing actions as classes enables powerful 
 * type checking in reducer functions.
 * 
 * See Discriminated Unions: https://www.typescriptlang.org/docs/handbook/advanced-types.html#discriminated-unions
 */
export class GetListSprintAction implements Action {
    type = ActionTypes.GET_LIST_SPRINT;
    payload;
    constructor() { }
}

export class GetListSprintSuccessAction implements Action {
    type = ActionTypes.GET_LIST_SPRINT_SUCCESS;

    constructor(public payload: Sprint[]) { }
}

export class GetListSprintFailAction implements Action {
    type = ActionTypes.GET_LIST_SPRINT_FAIL;

    constructor(public payload: any) { }
}

/**
 * Export a type alias of all actions in this action group
 * so that reducers can easily compose action types
 */
export type Actions
    = GetListSprintAction
    | GetListSprintSuccessAction
    | GetListSprintFailAction;