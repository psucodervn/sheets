interface TableColumn {
  name: string;
  label: string;
  field: string | Function;
  sortable?: boolean;
  align?: 'left' | 'right';
  sort?: Function;
  format?: Function;
}

interface TablePagination {
  sortBy?: string;
  descending?: boolean;
  page?: number;
  rowsPerPage?: number;
  rowsNumber?: number;
}

export {
  TableColumn,
  TablePagination,
};
