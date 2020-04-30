{{- if not .NoHooks -}}
{{- $alias := .Aliases.Table .Table.Name }}

import "github.com/rs/xid"

func init() {
  Add{{$alias.UpSingular}}Hook(boil.BeforeInsertHook, func(ctx context.Context, executor boil.ContextExecutor, {{$alias.DownSingular}} *{{$alias.UpSingular}}) error {
    if len({{$alias.DownSingular}}.ID) == 0 {
      {{$alias.DownSingular}}.ID = xid.New().String()
    }
    return nil
  })
  Add{{$alias.UpSingular}}Hook(boil.BeforeUpdateHook, func(ctx context.Context, executor boil.ContextExecutor, {{$alias.DownSingular}} *{{$alias.UpSingular}}) error {
    if len({{$alias.DownSingular}}.ID) == 0 {
      {{$alias.DownSingular}}.ID = xid.New().String()
    }
    return nil
  })
  Add{{$alias.UpSingular}}Hook(boil.BeforeUpsertHook, func(ctx context.Context, executor boil.ContextExecutor, {{$alias.DownSingular}} *{{$alias.UpSingular}}) error {
    if len({{$alias.DownSingular}}.ID) == 0 {
      {{$alias.DownSingular}}.ID = xid.New().String()
    }
    return nil
  })
}
{{- end}}
