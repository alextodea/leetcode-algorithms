package custom

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type Photo struct {
	rowIndex  int
	cityName  string
	timestamp time.Time
	extension string
}

type ProcessedPhoto struct {
	rowIndex int
	name     string
}

func Solution(S string) string {
	if len(S) < 1 {
		return S
	}

	var photosArr []*Photo

	for i, line := range strings.Split(strings.TrimSuffix(S, "\n"), "\n") {
		photo := parsePhoto(line, i)
		photosArr = append(photosArr, photo)
	}

	photosGrouppedByCity := groupPhotosByCity(photosArr)
	processedGrouppedPhotos := processGrouppedPhotos(photosGrouppedByCity)

	var output strings.Builder

	for _, photo := range photosArr {
		groupedByCityArr := processedGrouppedPhotos[photo.cityName]
		for _, processedPhoto := range groupedByCityArr {
			if processedPhoto.rowIndex == photo.rowIndex {
				output.WriteString(processedPhoto.name + "\n")
			}
		}
	}
	return output.String()
}

func processGrouppedPhotos(photosGrouppedByCity map[string][]*Photo) map[string][]*ProcessedPhoto {
	processedPhotosGrouppedByCity := make(map[string][]*ProcessedPhoto)
	for cityName, photos := range photosGrouppedByCity {
		processedPhotosGrouppedByCity[cityName] = processGroup(photos)
	}
	return processedPhotosGrouppedByCity
}
func processGroup(photos []*Photo) []*ProcessedPhoto {
	sort.SliceStable(photos, func(i, j int) bool {
		return photos[i].timestamp.Before(photos[j].timestamp)
	})

	var processedPhotos []*ProcessedPhoto

	for i, photo := range photos {
		photoOrder := strconv.Itoa(i + 1)
		processedPhotos = append(processedPhotos, &ProcessedPhoto{name: photo.cityName + photoOrder + "." + photo.extension, rowIndex: photo.rowIndex})
	}

	return processedPhotos
}

func groupPhotosByCity(photosArr []*Photo) map[string][]*Photo {
	var photosMap = map[string][]*Photo{}
	for _, photo := range photosArr {
		photosMap[photo.cityName] = append(photosMap[photo.cityName], photo)
	}
	return photosMap
}

func parsePhoto(line string, i int) *Photo {
	splittedLine := strings.Split(line, ", ")

	_, timestamp := convertStringToTimestamp(splittedLine[2])

	return &Photo{rowIndex: i, extension: strings.Split(splittedLine[0], ".")[1], timestamp: timestamp, cityName: splittedLine[1]}
}

// func addPhotoDetailsToNodeMappings(photoDetails *PhotoDetails) {

// 	_, nodeExistsInMapping := nodesMap[photoDetails.cityName]

// 	if !nodeExistsInMapping {
// 		nodesMap[photoDetails.cityName] = &PhotoNode{timestamp: photoDetails.timestamp, extension: photoDetails.extension, next: nil}
// 	}

// 	head := nodesMap[photoDetails.cityName]

// 	nodesMap[photoDetails.cityName] = createPhotoNode(head, photoDetails)

// 	a := nodesMap
// 	fmt.Print(a)
// }

func convertStringToTimestamp(s string) (error, time.Time) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, s)

	if err != nil {
		return err, t
	}

	return nil, t
}
