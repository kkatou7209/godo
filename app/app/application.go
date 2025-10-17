package app

import (
	"github.com/kkatou7209/godo/app/port/in/usecase"
	"github.com/kkatou7209/godo/app/port/out/password"
	"github.com/kkatou7209/godo/app/port/out/persistence"
	"github.com/kkatou7209/godo/app/service"
)

// Application configuration.
type Application struct {
	createTodoPersistence persistence.CreateTodoPersistence
	listTodoPersistence persistence.ListTodoPersistence
	updateTodoPersistence persistence.UpdateTodoPersistence
	deleteTodoPersistence persistence.DeleteTodoPersistence
	getTodoPersistence persistence.GetTodoPersistence
	getUserPersistence persistence.GetUserPersistence
	updateUserPersistence persistence.UpdateUserPersistence
	createUserPersistence persistence.CreateUserPersistence
	passwordHasher password.PasswordHasher
}

func New() *Application {
	return &Application{
		createTodoPersistence: nil,
		listTodoPersistence: nil,
		updateTodoPersistence: nil,
		deleteTodoPersistence: nil,
		getTodoPersistence: nil,
		getUserPersistence: nil,
		updateUserPersistence: nil,
		createUserPersistence: nil,
		passwordHasher: nil,
	}
}

func (a *Application) SetCreateTodoPersistence(createTodoPersistence persistence.CreateTodoPersistence) *Application {
	a.createTodoPersistence = createTodoPersistence
	return a
}

func (a *Application) SetUpdateTodoPersistence(updateTodoPersistence persistence.UpdateTodoPersistence) *Application {
	a.updateTodoPersistence = updateTodoPersistence
	return a
}

func (a *Application) SetGetTodoPersistence(getTodoPersistence persistence.GetTodoPersistence) *Application {
	a.getTodoPersistence = getTodoPersistence
	return a
}

func (a *Application) SetListTodoPersistence(listTodoPersistence persistence.ListTodoPersistence) *Application {
	a.listTodoPersistence = listTodoPersistence
	return a
}

func (a *Application) SetDeleteTodoPersistence(deleteTodoPersistence persistence.DeleteTodoPersistence) *Application {
	a.deleteTodoPersistence = deleteTodoPersistence
	return a
}

func (a *Application) SetGetUserPersistence(getUserPersistence persistence.GetUserPersistence) *Application {
	a.getUserPersistence = getUserPersistence
	return a
}

func (a *Application) SetUpdateUserPersistence(updateUserPersistence persistence.UpdateUserPersistence) *Application {
	a.updateUserPersistence = updateUserPersistence
	return a
}

func (a *Application) SetCreateUserPersistence(createUserPersistence persistence.CreateUserPersistence) *Application {
	a.createUserPersistence = createUserPersistence
	return a
}

func (a *Application) SetPasswordHasher(passwordHasher password.PasswordHasher) *Application {
	a.passwordHasher = passwordHasher
	return a
}

func (a *Application) AddUserUsecase() usecase.AddUserUsecase {
	return service.NewAddUserService(
		a.createUserPersistence,
		a.getUserPersistence,
		a.passwordHasher,
	)
}

func (a *Application) GetUserUsecase() usecase.GetUserUsecase {
	return service.NewGetUserService(a.getUserPersistence)
}

func (a *Application) ChangeUserInfoUsecase() usecase.ChangeUserInfoUsecase {
	return service.NewChangeUserInfoService(a.updateUserPersistence, a.getUserPersistence)
}

func (a *Application) ChangeUserPasswordUsecase() usecase.ChangeUserPasswordUsecase {
	return service.NewChangeUserPasswordService(a.updateUserPersistence, a.getUserPersistence, a.passwordHasher)
}

func (a *Application) LoginUsecase() usecase.LoginUsecase {
	return service.NewLoginService(a.getUserPersistence, a.passwordHasher)
}

func (a *Application) AddTodoUsecase() usecase.AddTodoUsecase {
	return service.NewAddTodoService(a.createTodoPersistence)
}

func (a *Application) GetTodoUsecase() usecase.GetTodoUsecase {
	return service.NewGetTodoService(a.getTodoPersistence)
}

func (a *Application) ListTodoUsecase() usecase.ListTodoUsecase {
	return service.NewListTodoService(a.listTodoPersistence)
}

func (a *Application) UpdateTodoUsecase() usecase.UpdateTodoUsecase {
	return service.NewUpdateTodoService(a.updateTodoPersistence, a.getTodoPersistence)
}

func (a *Application) CompleteTodoUsecase() usecase.CompleteTodoUsecase {
	return service.NewCompleteTodoService(a.updateTodoPersistence, a.getTodoPersistence)
}

func (a *Application) UncompleteTodoUsecase() usecase.UncompleteTodoUsecase {
	return service.NewUncompleteTodoService(a.updateTodoPersistence, a.getTodoPersistence)
}

func (a *Application) DeleteTodoUsecase() usecase.DeleteTodoUsecase {
	return service.NewDeleteTodoService(a.deleteTodoPersistence, a.getTodoPersistence)
}
