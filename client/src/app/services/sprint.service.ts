import 'rxjs/add/operator/map';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import { Sprint } from '../models/sprint';


@Injectable()
export class SprintService {
    private API_PATH: string = 'http://localhost:1323/sprints';

    constructor(private http: Http) { }

    getListSprints(): Observable<Sprint[]> {
        return this.http.get(`${this.API_PATH}`)
            .map(res => res.json() || []);
    }

}