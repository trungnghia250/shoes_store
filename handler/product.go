package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/trungnghia250/shoes_store/db"
	"github.com/trungnghia250/shoes_store/model"
	"github.com/trungnghia250/shoes_store/service"
	"math"
	"mime/multipart"

	"strconv"
)

func ListBrandProduct(c *fiber.Ctx) error {
	brandName := c.Params("brand_name")

	products, err := service.ListBrandProduct(c, brandName)
	if err != nil {
		return err
	}
	return c.JSON(products)
}

func ListProductBySize(c *fiber.Ctx) error {
	size := c.Params("size")
	sizeInt, _ := strconv.ParseInt(size, 10, 32)

	products, err := service.ListProductBySize(c, int32(sizeInt))
	if err != nil {
		return err
	}
	return c.JSON(products)
}

func GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, _ := strconv.ParseInt(id, 10, 32)

	product, err := service.GetProductByID(c, int32(idInt))
	if err != nil {
		return err
	}

	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	req := new(model.Product)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	image, _ := OpenFormFileSafe(c)
	if image != nil {
		storeId := fmt.Sprintf("test_upload")
		location, err := UploadImage(image, storeId)
		if err != nil {
			return err
		}
		req.Link = location
	}

	customer, err := service.CreateProduct(c, req)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}

func ListAllProduct(c *fiber.Ctx) error {
	products, err := service.ListAllProduct(c)
	if err != nil {
		return err
	}

	return c.JSON(products)
}

func OpenFormFileSafe(c *fiber.Ctx) (*multipart.FileHeader, error) {
	fh, err := c.FormFile("image")
	if err != nil {
		if err.Error() == "there is no uploaded file associated with the given key" {
			return nil, nil
		}

		return nil, err
	}

	return fh, nil
}

func UploadImage(fh *multipart.FileHeader, bucket string) (string, error) {
	body, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer func() {
		if err := body.Close(); err != nil {
			fmt.Sprintf("error here")
		}
	}()

	buffer, err := getImageBuffer(body)

	if !filetype.IsImage(buffer) {
		return "", fiber.NewError(fiber.StatusBadRequest, "only accept image upload")
	}

	uploader := s3manager.NewUploader(db.Sess)
	kind, err := filetype.Match(buffer)
	if err != nil {
		return "", err
	}

	filename := uuid.New().String()
	path := fmt.Sprintf("%s/%s.%s", bucket, filename, kind.Extension)
	resp, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("goshop-images"),
		ACL:         aws.String("public-read"),
		Key:         aws.String(path),
		ContentType: aws.String(kind.MIME.Value),
		Body:        body,
	})

	if err != nil {
		return "", err
	}

	return resp.Location, nil
}

func getImageBuffer(body multipart.File) ([]byte, error) {
	buffer := make([]byte, 512)
	_, err := body.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UpdateProduct(c *fiber.Ctx) error {
	req := new(model.UpdateRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	body := new(model.Product)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	err := service.UpdateProduct(c, req.Id, body)
	if err != nil {
		return err
	}

	return c.JSON(DefaultResponse{StatusCode: fiber.StatusOK})
}

func DeleteProduct(c *fiber.Ctx) error {
	req := new(model.GetCusRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	err := service.DeleteProductByID(c, req.ID)
	if err != nil {
		return err
	}
	return c.JSON(DefaultResponse{StatusCode: fiber.StatusOK})
}

func Rating(c *fiber.Ctx) error {
	req := new(model.UpdateRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}
	body := new(model.RatingRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}

	product, err := service.GetProductByID(c, req.Id)
	if err != nil {
		return err
	}

	product.Rating.Value = math.Floor((product.Rating.Value*float64(product.Rating.TotalRating)+float64(body.Value))/float64(product.Rating.TotalRating+1)/0.01) * 0.01
	product.Rating.TotalRating += 1

	err = service.UpdateProduct(c, req.Id, product)
	if err != nil {
		return err
	}
	return c.JSON(product.Rating)
}
