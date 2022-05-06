# -*- coding: utf-8 -*-
import hashlib

import flask_restful as restful

from ..validators import request_validate, response_filter


class Resource(restful.Resource):
    method_decorators = [request_validate, response_filter]


class Vator(object):

    def __init__(self, floors, car_ct=1):
        self.floor_list = floors

        self.first_floor = None
        self.floor_map = {}
        for ix in range(len(floors)):
            unhashed_floor = ('floor-%s' % ix).encode('utf-8')
            fid = hashlib.sha1(unhashed_floor).hexdigest()
            self.floor_map[fid] = floors[ix]
            if self.first_floor is None:
                self.first_floor = fid

        self.car_map = {}
        self.car_current_floor = {}
        for ix in range(car_ct):
            name = ('Car-%s' % ix).encode('utf-8')
            cid = hashlib.sha1(name).hexdigest()
            self.car_map[cid] = name
            self.car_current_floor[cid] = self.first_floor

    def floor_count(self):
        return len(self.floor_list)

    def inventory(self):
        results = []
        for fid, name in self.floor_map.iteritems():
            results.append({'id': fid, 'name': name})
        for fid, name in self.car_map.iteritems():
            results.append({'id': fid, 'name': name})
        return results

    def current_floor(self, car_id):
        newfloormap = self.floor_map
        floor_id = self.car_current_floor[car_id]
        return {'id': floor_id, 'name': self.floor_map[floor_id]}

    def find_closest_car(self, floor_id):
        cars = self.car_map
        floorid = {}
        for id, name in cars.items():
            currentfloor = self.current_floor(id)
            floorid[id] = currentfloor['name']
        pcurrentfloor = self.floor_map[floor_id]
        pcurrentfloor = self.floor_list.index(pcurrentfloor)

        floordistance = {}

        for key, values in floorid:
            floordistance[key] = abs(self.floor_list.index(value)-pcurrentfloor)
        minval = min(floordistance.values())

        for key, val in floordistance:
            if val == minval:
                return cars[key]


    def call_car(self, floor_id):
        findcar = self.find_closest_car(floor_id)
        return findcar


elevator = Vator(['B2', 'B1', 'MZ', 'F1', 'F2', 'F3', 'F4', 'F5', 'F6', 'F7'], 2)
