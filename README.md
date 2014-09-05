# orgstreak

Get back `json` representation of your organization's public members' public
GitHub contributions.

## Installation

```
$ go get github.com/fly/orgstreak
```

Or if you just want the binary, you can grab the one for your architecture of
choice from the [releases
page](https://github.com/fly/orgstreak/releases/tag/v0.0.3)

## Usage

```
$ orgstreak <orgName>
```

i.e.:

```
$ orgstreak linode
[["2014-02-24",4],["2014-02-26",8],["2014-04-02",11],["2014-05-09",9],["2014-06-08",21],["2014-08-28",27],["2013-09-05",3],["2014-01-21",6],["2014-05-25",14],["2014-06-06",14],["2014-07-11",34],["2014-08-09",37],["2014-01-12",9],["2014-05-16",17],["2014-02-10",11],["2014-03-23",35],["2014-05-20",40],["2014-05-26",16],["2014-06-17",24],["2014-08-15",48],["2013-09-30",30],["2014-01-02",11],["2013-11-28",7],["2013-12-15",13],["2014-08-02",46],["2013-11-07",18],["2014-06-23",16],["2013-12-23",13],["2014-02-15",86],["2014-07-24",21],["2013-09-20",4],["2013-10-22",5],["2014-02-22",57],["2014-03-06",16],["2014-04-12",23],["2014-06-03",50],["2014-08-25",46],["2013-10-04",10],["2013-11-22",33],["2013-12-13",11],["2014-02-02",28],["2014-05-14",24],["2013-11-11",9],["2013-11-24",12],["2014-05-22",10],["2014-05-24",15],["2014-08-29",14],["2013-10-26",16],["2014-03-13",49],["2014-04-10",29],["2014-05-05",34],["2014-06-28",17],["2014-07-03",13],["2014-08-13",10],["2013-11-29",24],["2014-02-08",8],["2014-02-25",7],["2014-08-11",34],["2013-11-03",26],["2013-11-09",4],["2014-01-09",14],["2014-06-13",18],["2014-06-26",49],["2014-07-12",39],["2014-07-15",21],["2014-08-31",77],["2013-09-18",21],["2013-11-16",6],["2014-06-21",30],["2014-08-05",35],["2014-01-19",5],["2014-02-13",10],["2013-12-10",2],["2013-12-27",28],["2014-03-28",19],["2014-04-21",14],["2014-06-22",21],["2014-07-05",37],["2013-10-18",6],["2013-10-21",6],["2014-01-04",10],["2014-03-02",27],["2014-04-14",20],["2014-06-02",79],["2014-07-17",7],["2014-07-22",13],["2013-11-27",16],["2013-12-12",6],["2014-07-30",35],["2014-08-07",18],["2014-01-30",15],["2014-02-23",81],["2014-02-28",19],["2014-05-07",15],["2014-05-30",17],["2014-07-04",29],["2013-10-30",5],["2013-12-18",21],["2014-06-15",25],["2014-01-14",3],["2014-06-04",27],["2014-02-16",62],["2014-06-05",25],["2014-07-20",27],["2014-07-29",32],["2013-10-02",7],["2013-10-05",16],["2014-02-20",16],["2014-04-30",22],["2014-07-27",67],["2014-08-03",27],["2013-09-23",14],["2013-12-31",6],["2014-03-09",3],["2014-05-17",55],["2014-07-09",24],["2013-10-19",6],["2014-01-29",10],["2013-11-17",13],["2014-01-16",5],["2014-01-25",9],["2014-02-05",13],["2014-08-06",40],["2014-08-18",36],["2013-12-28",14],["2014-01-17",7],["2013-11-30",14],["2014-01-26",61],["2014-06-25",34],["2014-07-01",25],["2013-09-03",6],["2013-11-21",51],["2014-03-03",10],["2014-05-08",13],["2014-06-10",46],["2014-06-29",74],["2013-12-09",5],["2014-01-13",8],["2014-03-17",23],["2014-03-27",22],["2014-05-06",9],["2013-11-05",22],["2014-03-10",28],["2013-10-24",11],["2014-04-29",36],["2014-07-08",25],["2014-07-26",50],["2013-09-08",16],["2013-09-11",9],["2013-12-11",9],["2014-01-01",5],["2014-03-11",19],["2014-03-22",8],["2014-08-26",19],["2013-10-08",7],["2013-11-23",54],["2014-07-31",33],["2014-08-14",15],["2013-11-26",5],["2013-12-04",19],["2013-10-16",4],["2013-11-06",20],["2014-01-23",3],["2014-07-28",25],["2013-09-10",11],["2013-09-26",4],["2014-06-09",31],["2014-06-19",12],["2013-09-15",20],["2013-10-09",8],["2014-01-10",4],["2014-02-04",12],["2014-03-01",44],["2014-03-31",18],["2014-06-01",14],["2014-08-12",22],["2013-09-21",9],["2014-01-07",12],["2014-01-24",29],["2014-03-18",17],["2014-03-24",17],["2014-06-11",28],["2014-07-19",15],["2013-10-20",22],["2013-11-18",21],["2014-03-05",26],["2014-04-09",43],["2014-07-13",16],["2013-10-23",20],["2013-12-05",3],["2013-10-01",6],["2013-10-15",9],["2013-11-04",24],["2014-01-06",2],["2014-03-07",7],["2014-05-01",14],["2013-09-12",2],["2013-09-24",43],["2014-05-31",15],["2013-11-20",18],["2013-12-01",14],["2014-03-26",53],["2014-07-18",11],["2013-09-01",4],["2013-11-19",89],["2014-05-03",29],["2014-05-15",7],["2014-06-07",21],["2013-09-17",11],["2014-02-27",17],["2014-01-03",3],["2014-02-06",27],["2014-03-15",50],["2014-04-06",47],["2014-04-16",51],["2014-07-02",59],["2013-10-25",17],["2013-12-16",5],["2014-08-10",23],["2014-08-24",35],["2013-11-02",19],["2013-11-12",9],["2014-06-24",9],["2013-09-02",7],["2013-09-22",19],["2014-03-20",31],["2014-05-04",16],["2013-09-07",5],["2014-02-19",15],["2013-11-14",12],["2013-11-25",22],["2014-01-28",5],["2014-02-14",12],["2014-07-14",18],["2014-08-20",38],["2013-10-07",8],["2013-10-11",8],["2013-10-14",9],["2013-11-13",9],["2014-01-15",8],["2014-03-19",16],["2014-05-02",55],["2014-05-12",23],["2014-05-23",30],["2013-12-14",2],["2013-12-24",17],["2013-12-29",11],["2014-02-01",17],["2014-02-18",22],["2014-07-07",26],["2014-07-25",22],["2014-08-08",21],["2014-08-27",53],["2013-10-13",17],["2013-12-03",3],["2014-01-11",13],["2014-03-04",25],["2014-04-22",15],["2014-08-21",20],["2013-09-19",7],["2013-10-03",9],["2014-02-07",37],["2014-05-27",8],["2014-06-16",19],["2014-08-22",58],["2013-09-28",9],["2013-12-06",2],["2013-12-26",5],["2014-02-03",29],["2014-02-11",19],["2014-04-17",10],["2014-04-27",23],["2014-07-06",32],["2013-09-06",14],["2013-09-27",9],["2014-04-11",14],["2014-04-13",36],["2014-06-20",20],["2014-06-30",64],["2014-01-08",7],["2014-02-21",22],["2013-11-08",15],["2014-01-18",2],["2014-04-03",15],["2014-08-30",94],["2013-09-09",11],["2013-10-28",12],["2013-12-08",3],["2013-12-30",6],["2014-04-05",19],["2014-04-08",13],["2014-06-14",46],["2013-11-10",4],["2014-03-30",19],["2014-04-19",25],["2014-07-10",15],["2014-03-08",32],["2014-04-04",15],["2014-08-01",21],["2014-08-04",16],["2014-09-01",7],["2013-12-22",4],["2014-04-24",5],["2013-10-10",8],["2013-12-25",33],["2014-01-20",8],["2014-02-12",23],["2014-02-17",31],["2014-04-01",16],["2013-08-31",6],["2013-09-29",35],["2014-06-12",23],["2014-07-21",9],["2013-09-25",10],["2013-10-06",29],["2013-10-27",11],["2014-01-05",6],["2014-03-12",31],["2014-03-21",64],["2013-09-13",2],["2013-09-16",7],["2014-05-21",10],["2014-01-22",13],["2014-01-27",6],["2014-04-26",11],["2014-05-19",27],["2013-12-19",2],["2013-12-20",11],["2014-04-20",58],["2014-07-23",11],["2013-09-14",5],["2013-10-29",4],["2014-03-16",23],["2014-04-23",23],["2014-05-11",15],["2013-09-04",6],["2013-10-31",8],["2014-04-15",24],["2014-04-18",11],["2014-04-25",52],["2014-05-28",8],["2014-05-29",14],["2014-06-18",9],["2013-12-17",16],["2014-03-29",28],["2014-08-23",13],["2014-05-13",20],["2014-05-18",45],["2013-12-21",16],["2014-01-31",11],["2014-02-09",21],["2014-03-14",28],["2014-03-25",52],["2014-04-07",10],["2013-10-12",26],["2013-11-15",10],["2014-04-28",25],["2014-05-10",11],["2014-06-27",16],["2014-07-16",32],["2013-10-17",5],["2013-12-07",13],["2014-08-16",59],["2014-08-17",31],["2013-11-01",14],["2013-12-02",13],["2014-08-19",40]]
```
