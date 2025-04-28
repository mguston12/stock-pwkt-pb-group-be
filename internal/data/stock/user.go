package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"
)

func (d Data) GetUserByUsername(ctx context.Context, username string) (stock.User, error) {
	user := stock.User{}

	if err := d.stmt[getUserByUsername].QueryRowxContext(ctx, username).StructScan(&user); err != nil {
		return user, errors.Wrap(err, "[DATA][GetUserByUsername]")
	}

	return user, nil
}

func (d Data) CreateUser(ctx context.Context, user stock.User) error {
	_, err := d.stmt[createUser].ExecContext(ctx,
		user.Username,
		user.Password,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateUser]")
	}
	return nil
}

func (d Data) UpdateUser(ctx context.Context, user stock.User) error {
	_, err := d.stmt[updateUser].ExecContext(ctx,
		user.Password,
		user.Username,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateUser]")
	}
	return nil
}

func (d Data) DeleteUser(ctx context.Context, username string) error {
	_, err := d.stmt[deleteUser].ExecContext(ctx, username)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteUser]")
	}
	return nil
}
