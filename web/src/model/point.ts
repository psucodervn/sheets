export interface IIssue {
  id: string;
  key: string;
  summary: string;
  point: number;
}

export class Issue implements IIssue {
  id: string;
  key: string;
  point: number;
  summary: string;

  constructor(id: string, key: string, point: number, summary: string) {
    this.id = id;
    this.key = key;
    this.point = point;
    this.summary = summary;
  }
}

export interface IUserPoint {
  name: string;
  displayName: string;
  pointTotal: number;
}

export class UserPoint implements IUserPoint {
  displayName: string;
  name: string;
  pointTotal: number;

  constructor(displayName: string, name: string, pointTotal: number) {
    this.displayName = displayName;
    this.name = name;
    this.pointTotal = pointTotal;
  }
}
