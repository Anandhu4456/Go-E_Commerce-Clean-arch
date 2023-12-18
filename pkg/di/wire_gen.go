// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/db"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository"
	"github.com/Anandhu4456/go-Ecommerce/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	gormDB, err := db.ConnectDB(cfg)
	if err != nil {
		return nil, err
	}
	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)
	inventoryRespository := repository.NewInventoryRepository(gormDB)
	inventoryUsecase := usecase.NewInventoryUsecase(inventoryRespository)
	inventoryHandler := handlers.NewInventoryHandler(inventoryUsecase)
	userRepository := repository.NewUserRepository(gormDB)
	offerRepository := repository.NewOfferRepository(gormDB)
	walletRepository := repository.NewWalletRepository(gormDB)
	userUsecase := usecase.NewUserUsecase(userRepository, offerRepository, walletRepository)
	userHandler := handlers.NewUserHandler(userUsecase)
	otpRepository := repository.NewOtpRepository(gormDB)
	otpUsecase := usecase.NewOtpUsecase(cfg, otpRepository)
	otpHandler := handlers.NewOtpHandler(otpUsecase)
	adminRepository := repository.NewAdminRepository(gormDB)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminHandler := handlers.NewAdminHandler(adminUsecase)
	cartRepository := repository.NewCartRepository(gormDB)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository, userRepository)
	cartUsecase := usecase.NewCartUsecase(cartRepository, inventoryRespository, userUsecase, paymentUsecase)
	cartHandler := handlers.NewCartHandler(cartUsecase)
	orderRepository := repository.NewOrderRepository(gormDB)
	couponRepository := repository.NewCouponRepository(gormDB)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, userUsecase, walletRepository, couponRepository)
	orderHandler := handlers.NewOrderHandler(orderUsecase)
	paymentHandler := handlers.NewPaymentHandler(paymentUsecase)
	wishlistRepository := repository.NewWishlistRepository(gormDB)
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepository)
	wishlistHandler := handlers.NewWishlistHandler(wishlistUsecase)
	offerUsecase := usecase.NewOfferUsecase(offerRepository)
	offerHandler := handlers.NewOfferHandler(offerUsecase)
	couponUsecase := usecase.NewCouponUsecase(couponRepository)
	couponHandler := handlers.NewCouponHandler(couponUsecase)
	serverHTTP := api.NewServerHttp(categoryHandler, inventoryHandler, userHandler, otpHandler, adminHandler, cartHandler, orderHandler, paymentHandler, wishlistHandler, offerHandler, couponHandler)
	return serverHTTP, nil
}
