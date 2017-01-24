export class Sprint {
    id : string;
	created : string;
	name : string;
	startDate : Date;
    endDate : Date;
    stream : Stream;
    capacity : number;
    velocity : number;
    planning : PlannedDay[];
    events : Event[];
}

export class Stream {
    name : string;
	boardId : string;
	laneId : string;
}

export class PlannedDay {
    day : Date;
	dayCapacity : number;
	theoricalPointToDo : number;
}

export class Event {
    eventDate : Date;
	name : string;
	duration : number;
}