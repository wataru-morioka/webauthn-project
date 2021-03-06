// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testPhotographs(t *testing.T) {
	t.Parallel()

	query := Photographs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPhotographsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhotographsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Photographs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhotographsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PhotographSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPhotographsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PhotographExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Photograph exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PhotographExists to return true, but got false.")
	}
}

func testPhotographsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	photographFound, err := FindPhotograph(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if photographFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPhotographsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Photographs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPhotographsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Photographs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPhotographsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	photographOne := &Photograph{}
	photographTwo := &Photograph{}
	if err = randomize.Struct(seed, photographOne, photographDBTypes, false, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}
	if err = randomize.Struct(seed, photographTwo, photographDBTypes, false, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = photographOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = photographTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Photographs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPhotographsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	photographOne := &Photograph{}
	photographTwo := &Photograph{}
	if err = randomize.Struct(seed, photographOne, photographDBTypes, false, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}
	if err = randomize.Struct(seed, photographTwo, photographDBTypes, false, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = photographOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = photographTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func photographBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func photographAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Photograph) error {
	*o = Photograph{}
	return nil
}

func testPhotographsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Photograph{}
	o := &Photograph{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, photographDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Photograph object: %s", err)
	}

	AddPhotographHook(boil.BeforeInsertHook, photographBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	photographBeforeInsertHooks = []PhotographHook{}

	AddPhotographHook(boil.AfterInsertHook, photographAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	photographAfterInsertHooks = []PhotographHook{}

	AddPhotographHook(boil.AfterSelectHook, photographAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	photographAfterSelectHooks = []PhotographHook{}

	AddPhotographHook(boil.BeforeUpdateHook, photographBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	photographBeforeUpdateHooks = []PhotographHook{}

	AddPhotographHook(boil.AfterUpdateHook, photographAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	photographAfterUpdateHooks = []PhotographHook{}

	AddPhotographHook(boil.BeforeDeleteHook, photographBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	photographBeforeDeleteHooks = []PhotographHook{}

	AddPhotographHook(boil.AfterDeleteHook, photographAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	photographAfterDeleteHooks = []PhotographHook{}

	AddPhotographHook(boil.BeforeUpsertHook, photographBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	photographBeforeUpsertHooks = []PhotographHook{}

	AddPhotographHook(boil.AfterUpsertHook, photographAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	photographAfterUpsertHooks = []PhotographHook{}
}

func testPhotographsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhotographsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(photographColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPhotographsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPhotographsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PhotographSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPhotographsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Photographs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	photographDBTypes = map[string]string{`ID`: `integer`, `SubTitle`: `character varying`, `Title`: `character varying`, `Mimetype`: `character varying`, `FileName`: `character varying`, `Size`: `integer`, `Data`: `bytea`, `CreatedDatetime`: `timestamp without time zone`, `ModifiedDatetime`: `timestamp without time zone`}
	_                 = bytes.MinRead
)

func testPhotographsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(photographPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(photographAllColumns) == len(photographPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, photographDBTypes, true, photographPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPhotographsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(photographAllColumns) == len(photographPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Photograph{}
	if err = randomize.Struct(seed, o, photographDBTypes, true, photographColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, photographDBTypes, true, photographPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(photographAllColumns, photographPrimaryKeyColumns) {
		fields = photographAllColumns
	} else {
		fields = strmangle.SetComplement(
			photographAllColumns,
			photographPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := PhotographSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPhotographsUpsert(t *testing.T) {
	t.Parallel()

	if len(photographAllColumns) == len(photographPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Photograph{}
	if err = randomize.Struct(seed, &o, photographDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Photograph: %s", err)
	}

	count, err := Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, photographDBTypes, false, photographPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Photograph struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Photograph: %s", err)
	}

	count, err = Photographs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
