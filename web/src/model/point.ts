export interface IIssue {
  id: string;
  key: string;
  summary: string;
  point: number;
  status: string;
}

export interface IUserPoint {
  name: string;
  displayName: string;
  pointTotal: number;
  wakatimeSeconds: number;
  wakatimeHuman: string;
  issues: IIssue[];
}
