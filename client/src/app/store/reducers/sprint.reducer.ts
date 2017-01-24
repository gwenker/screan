import { createSelector } from 'reselect';
import { Sprint } from '../../models/sprint';
import * as sprintAction from '../actions/sprint.action';


export interface State {
    sprints: { [id: string]: Sprint };
};

const initialState: State = {
    sprints: {}
};

export function reducer(state = initialState, action: sprintAction.Actions): State {
    switch (action.type) {
        // Sprint list success
        case sprintAction.ActionTypes.GET_LIST_SPRINT_SUCCESS: {
            const sprintsResult = action.payload;
            const sprintsMap = sprintsResult.reduce((sprints: { [id: string]: Sprint }, sprint: Sprint) => {
                return Object.assign(sprints, {
                    [sprint.id]: sprint
                });
            }, {});

            return {
                sprints: sprintsMap
            };
        }
        default: {
            return state;
        }
    }
}

/**
 * Because the data structure is defined within the reducer it is optimal to
 * locate our selector functions at this level. If store is to be thought of
 * as a database, and reducers the tables, selectors can be considered the
 * queries into said database. Remember to keep your selectors small and
 * focused so they can be combined and composed to fit each particular
 * use-case.
 */

export const getSprints = (state: State) => state.sprints;
