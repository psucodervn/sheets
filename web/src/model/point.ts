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
  issues: Issue[];
}

export class UserPoint implements IUserPoint {
  displayName: string;
  name: string;
  pointTotal: number;
  issues: Issue[];

  constructor(displayName: string, name: string, pointTotal: number, issues: Issue[]) {
    this.displayName = displayName;
    this.name = name;
    this.pointTotal = pointTotal;
    this.issues = issues;
  }
}
