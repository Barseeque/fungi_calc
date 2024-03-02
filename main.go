package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math"
	"strconv"
)

func main() {

	versa := "build: 020324 14:45"
	myApp := app.New()
	myWindow := myApp.NewWindow("Калькулятор компонентов субстрата")
	myWindow.Resize(fyne.NewSize(600, 360))

	totalwater := widget.NewEntry()
	totalwater.Resize(fyne.NewSize(55, 35))
	totalwater.Move(fyne.NewPos(520, 290))

	totalmass := widget.NewEntry()
	totalmass.Text = "1000"
	totalmass.Resize(fyne.NewSize(55, 35))
	totalmass.Move(fyne.NewPos(210, 250))

	Humidity := widget.NewEntry()
	Humidity.Text = "65"
	Humidity.Resize(fyne.NewSize(55, 35))
	Humidity.Move(fyne.NewPos(210, 290))

	addWaterZag := widget.NewLabel("Добавьте воды, г")
	addWaterZag.Resize(fyne.NewSize(250, 35))
	addWaterZag.Move(fyne.NewPos(390, 290))

	totalmassZag := widget.NewLabel("Масса готового субстрата, г")
	totalmassZag.Resize(fyne.NewSize(200, 35))
	totalmassZag.Move(fyne.NewPos(0, 250))

	HumidityZag := widget.NewLabel("Влажность гот. субстрата, %")

	HumidityZag.Resize(fyne.NewSize(200, 35))
	HumidityZag.Move(fyne.NewPos(0, 290))

	zagol := widget.NewLabel("Компонент                                  Влажность (%) 	Содержание             Вес (г)      Воды (г)")
	zagol.Move(fyne.NewPos(0, 0))
	zagol1 := widget.NewLabel("в субстрате (%)")
	zagol1.Move(fyne.NewPos(320, 15))

	variants := []string{
		"Берёзовый RUF",
		"Буковые пеллеты",
		"Пеллеты из лузги",
		"Отруби",
		"Пеллеты из люцерны",
		"Пеллеты комбикорм",
		"Мел",
		"_________________",
	}

	sod1 := widget.NewEntry()
	sod1.Text = "20"
	sod1.Resize(fyne.NewSize(55, 35))
	sod1.Move(fyne.NewPos(325, 50))

	proc1 := widget.NewEntry()
	proc1.Text = "0"
	proc1.Resize(fyne.NewSize(55, 35))
	proc1.Move(fyne.NewPos(210, 50))

	ves1 := widget.NewEntry()
	ves1.Resize(fyne.NewSize(55, 35))
	ves1.Move(fyne.NewPos(460, 50))

	komp1 := widget.NewSelect(variants,
		func(s string) {
			a, b := getHumid(s)
			proc1.Text = a
			sod1.Text = b
			proc1.Refresh()
			sod1.Refresh()
		},
	)

	komp1.Resize(fyne.NewSize(200, 35))
	komp1.Move(fyne.NewPos(5, 50))
	komp1.SetSelected("Берёзовый RUF")
	komp1.Refresh()

	wat1 := widget.NewEntry()
	wat1.Resize(fyne.NewSize(55, 35))
	wat1.Move(fyne.NewPos(520, 50))

	proc2 := widget.NewEntry()
	proc2.Text = "0"
	proc2.Resize(fyne.NewSize(55, 35))
	proc2.Move(fyne.NewPos(210, 90))

	ves2 := widget.NewEntry()
	//	ves2.Text = "100"
	ves2.Resize(fyne.NewSize(55, 35))
	ves2.Move(fyne.NewPos(460, 90))

	sod2 := widget.NewEntry()
	sod2.Text = "20"
	sod2.Resize(fyne.NewSize(55, 35))
	sod2.Move(fyne.NewPos(325, 90))

	komp2 := widget.NewSelect(variants,
		func(s string) {
			a, b := getHumid(s)

			proc2.Text = a
			sod2.Text = b
			proc2.Refresh()
			sod2.Refresh()

		},
	)
	komp2.Resize(fyne.NewSize(200, 35))
	komp2.Move(fyne.NewPos(5, 90))
	komp2.SetSelected("_________________")
	komp2.Refresh()

	wat2 := widget.NewEntry()
	wat2.Resize(fyne.NewSize(55, 35))
	wat2.Move(fyne.NewPos(520, 90))

	proc3 := widget.NewEntry()
	proc3.Text = "0"
	proc3.Resize(fyne.NewSize(55, 35))
	proc3.Move(fyne.NewPos(210, 130))

	ves3 := widget.NewEntry()
	//	ves3.Text = "100"
	ves3.Resize(fyne.NewSize(55, 35))
	ves3.Move(fyne.NewPos(460, 130))

	sod3 := widget.NewEntry()
	sod3.Text = "20"
	sod3.Resize(fyne.NewSize(55, 35))
	sod3.Move(fyne.NewPos(325, 130))

	komp3 := widget.NewSelect(variants,
		func(s string) {
			a, b := getHumid(s)
			proc3.Text = a
			proc3.Refresh()
			sod3.Text = b
			sod3.Refresh()
		},
	)
	komp3.Resize(fyne.NewSize(200, 35))
	komp3.Move(fyne.NewPos(5, 130))
	komp3.SetSelected("_________________")
	komp3.Refresh()

	wat3 := widget.NewEntry()
	wat3.Resize(fyne.NewSize(55, 35))
	wat3.Move(fyne.NewPos(520, 130))

	proc4 := widget.NewEntry()
	proc4.Text = "0"
	proc4.Resize(fyne.NewSize(55, 35))
	proc4.Move(fyne.NewPos(210, 170))

	ves4 := widget.NewEntry()
	//	ves4.Text = "100"
	ves4.Resize(fyne.NewSize(55, 35))
	ves4.Move(fyne.NewPos(460, 170))

	sod4 := widget.NewEntry()
	sod4.Text = "20"
	sod4.Resize(fyne.NewSize(55, 35))
	sod4.Move(fyne.NewPos(325, 170))

	komp4 := widget.NewSelect(variants,
		func(s string) {
			a, b := getHumid(s)
			proc4.Text = a
			proc4.Refresh()
			sod4.Text = b
			sod4.Refresh()
		},
	)
	komp4.Resize(fyne.NewSize(200, 35))
	komp4.Move(fyne.NewPos(5, 170))
	komp4.SetSelected("_________________")
	komp4.Refresh()

	wat4 := widget.NewEntry()
	wat4.Resize(fyne.NewSize(55, 35))
	wat4.Move(fyne.NewPos(520, 170))

	proc5 := widget.NewEntry()
	proc5.Text = "0"
	proc5.Resize(fyne.NewSize(55, 35))
	proc5.Move(fyne.NewPos(210, 210))

	ves5 := widget.NewEntry()
	//	ves5.Text = "100"
	ves5.Resize(fyne.NewSize(55, 35))
	ves5.Move(fyne.NewPos(460, 210))

	sod5 := widget.NewEntry()
	sod5.Text = "20"
	sod5.Resize(fyne.NewSize(55, 35))
	sod5.Move(fyne.NewPos(325, 210))

	komp5 := widget.NewSelect(variants,
		func(s string) {
			a, b := getHumid(s)
			proc5.Text = a
			proc5.Refresh()
			sod5.Text = b
			sod5.Refresh()

		},
	)
	komp5.Resize(fyne.NewSize(200, 35))
	komp5.Move(fyne.NewPos(5, 210))
	komp5.SetSelected("_________________")
	komp5.Refresh()

	wat5 := widget.NewEntry()
	wat5.Resize(fyne.NewSize(55, 35))
	wat5.Move(fyne.NewPos(520, 210))

	telega := widget.NewLabel("t.me/gray_monsta")
	telega.Resize(fyne.NewSize(200, 25))
	telega.Move(fyne.NewPos(460, 320))

	version := widget.NewLabel(versa)
	version.Resize(fyne.NewSize(200, 25))
	version.Move(fyne.NewPos(0, 320))

	// Создаем кнопку для расчета
	calcByProc := widget.NewButton("Счёт", func() {

		fullweight, _ := strconv.Atoi(totalmass.Text)
		humidPercent, _ := strconv.Atoi(Humidity.Text)
		soder1, _ := strconv.Atoi(sod1.Text) //содержание компонента в субстрате
		soder2, _ := strconv.Atoi(sod2.Text) //содержание компонента в субстрате
		soder3, _ := strconv.Atoi(sod3.Text) //содержание компонента в субстрате
		soder4, _ := strconv.Atoi(sod4.Text) //содержание компонента в субстрате
		soder5, _ := strconv.Atoi(sod5.Text) //содержание компонента в субстрате

		if soder1+soder2+soder3+soder4+soder5 < 100 {
			if soder5 != 0 {
				soder5 = 100 - soder1 - soder2 - soder3 - soder4

			} else if soder4 != 0 {
				soder4 = 100 - soder1 - soder2 - soder3

			} else if soder2 != 0 {
				soder3 = 100 - soder1 - soder2
			} else {
				soder1 = 100
			}
		}

		if soder1 > 100 {
			soder1 = 100
			soder2, soder3, soder4, soder5 = 0, 0, 0, 0

		} else if (soder1 + soder2) > 100 {
			soder2 = 100 - soder1
			soder3, soder4, soder5 = 0, 0, 0

		} else if (soder1 + soder2 + soder3) > 100 {
			soder3 = 100 - soder1 - soder2
			soder4, soder5 = 0, 0

		} else if (soder1 + soder2 + soder3 + soder4) > 100 {
			soder4 = 100 - soder1 - soder2 - soder3
			soder5 = 0

		} else {
			soder5 = 100 - soder1 - soder2 - soder3 - soder4
			sod5.Text = strconv.Itoa(soder5)
			sod5.Refresh()

		}

		sod1.Text = strconv.Itoa(soder1)
		sod1.Refresh()
		sod2.Text = strconv.Itoa(soder2)
		sod2.Refresh()
		sod3.Text = strconv.Itoa(soder3)
		sod3.Refresh()
		sod4.Text = strconv.Itoa(soder4)
		sod4.Refresh()
		sod5.Text = strconv.Itoa(soder5)
		sod5.Refresh()

		wta := 0
		procent1, _ := strconv.Atoi(proc1.Text) //процент влажности для данного компонента
		normalMassComponent, waterToAdd := calculation(fullweight, humidPercent, soder1, procent1)
		wta += waterToAdd
		ves1.Text = strconv.Itoa(normalMassComponent)
		ves1.Refresh()
		wat1.Text = strconv.Itoa(waterToAdd)
		wat1.Refresh()

		procent2, _ := strconv.Atoi(proc2.Text) //процент влажности для данного компонента
		normalMassComponent, waterToAdd = calculation(fullweight, humidPercent, soder2, procent2)
		wta += waterToAdd
		ves2.Text = strconv.Itoa(normalMassComponent)
		ves2.Refresh()
		wat2.Text = strconv.Itoa(waterToAdd)
		wat2.Refresh()

		procent3, _ := strconv.Atoi(proc3.Text) //процент влажности для данного компонента
		normalMassComponent, waterToAdd = calculation(fullweight, humidPercent, soder3, procent3)
		wta += waterToAdd
		ves3.Text = strconv.Itoa(normalMassComponent)
		ves3.Refresh()
		wat3.Text = strconv.Itoa(waterToAdd)
		wat3.Refresh()

		procent4, _ := strconv.Atoi(proc4.Text) //процент влажности для данного компонента
		normalMassComponent, waterToAdd = calculation(fullweight, humidPercent, soder4, procent4)
		wta += waterToAdd
		ves4.Text = strconv.Itoa(normalMassComponent)
		ves4.Refresh()
		wat4.Text = strconv.Itoa(waterToAdd)
		wat4.Refresh()

		procent5, _ := strconv.Atoi(proc5.Text) //процент влажности для данного компонента
		normalMassComponent, waterToAdd = calculation(fullweight, humidPercent, soder5, procent5)
		wta += waterToAdd
		ves5.Text = strconv.Itoa(normalMassComponent)
		ves5.Refresh()
		wat5.Text = strconv.Itoa(waterToAdd)
		wat5.Refresh()

		totalwater.Text = strconv.Itoa(wta)
		totalwater.Refresh()

	})

	calcByProc.Move(fyne.NewPos(325, 250))
	calcByProc.Resize(fyne.NewSize(55, 35))

	calcByWeight := widget.NewButton("Счёт", func() {

		humidPercent, _ := strconv.ParseFloat(Humidity.Text, 64)

		weight1, _ := strconv.ParseFloat(ves1.Text, 64) //содержание компонента в субстрате
		weight2, _ := strconv.ParseFloat(ves2.Text, 64) //содержание компонента в субстрате
		weight3, _ := strconv.ParseFloat(ves3.Text, 64) //содержание компонента в субстрате
		weight4, _ := strconv.ParseFloat(ves4.Text, 64) //содержание компонента в субстрате
		weight5, _ := strconv.ParseFloat(ves5.Text, 64) //содержание компонента в субстрате
		//
		if weight1 == 0 {
			weight1 = 1000
			ves1.Text = "1000"
			ves1.Refresh()
		}

		var procent1, procent2, procent3, procent4, procent5, haveWater, haveDry, waterToAdd, wta, totmass float64
		wta, totmass = 0, 0

		procent1, _ = strconv.ParseFloat(proc1.Text, 64)   //процент влажности для данного компонента
		haveWater = weight1 * procent1 / 100               //сколько воды уже есть в компоненте
		haveDry = weight1 - weight1*procent1/100           //сухая масса субстрата
		wetWeight1 := haveDry * 100 / (100 - humidPercent) //масса размоченного субстрата
		waterToAdd = wetWeight1 - haveDry - haveWater
		wta += waterToAdd
		totmass += wetWeight1
		wat1.Text = strconv.Itoa(int(waterToAdd))
		wat1.Refresh()

		procent2, _ = strconv.ParseFloat(proc2.Text, 64)   //процент влажности для данного компонента
		haveWater = weight2 * procent2 / 100               //сколько воды уже есть в компоненте
		haveDry = weight2 - weight2*procent2/100           //сухая масса субстрата
		wetWeight2 := haveDry * 100 / (100 - humidPercent) //масса размоченного субстрата
		waterToAdd = wetWeight2 - haveDry - haveWater
		wta += waterToAdd
		totmass += wetWeight2
		wat2.Text = strconv.Itoa(int(waterToAdd))
		wat2.Refresh()

		procent3, _ = strconv.ParseFloat(proc3.Text, 64)   //процент влажности для данного компонента
		haveWater = weight3 * procent3 / 100               //сколько воды уже есть в компоненте
		haveDry = weight3 - weight3*procent3/100           //сухая масса субстрата
		wetWeight3 := haveDry * 100 / (100 - humidPercent) //масса размоченного субстрата
		waterToAdd = wetWeight3 - haveDry - haveWater
		wta += waterToAdd
		totmass += wetWeight3
		wat3.Text = strconv.Itoa(int(waterToAdd))
		wat3.Refresh()

		procent4, _ = strconv.ParseFloat(proc4.Text, 64)   //процент влажности для данного компонента
		haveWater = weight4 * procent4 / 100               //сколько воды уже есть в компоненте
		haveDry = weight4 - weight4*procent4/100           //сухая масса субстрата
		wetWeight4 := haveDry * 100 / (100 - humidPercent) //масса размоченного субстрата
		waterToAdd = wetWeight4 - haveDry - haveWater
		wta += waterToAdd
		totmass += wetWeight4
		wat4.Text = strconv.Itoa(int(waterToAdd))
		wat4.Refresh()

		procent5, _ = strconv.ParseFloat(proc5.Text, 64)   //процент влажности для данного компонента
		haveWater = weight5 * procent5 / 100               //сколько воды уже есть в компоненте
		haveDry = weight5 - weight5*procent5/100           //сухая масса субстрата
		wetWeight5 := haveDry * 100 / (100 - humidPercent) //масса размоченного субстрата
		waterToAdd = wetWeight5 - haveDry - haveWater
		wta += waterToAdd
		totmass += wetWeight5
		wat5.Text = strconv.Itoa(int(waterToAdd))
		wat5.Refresh()

		sodr1 := int(wetWeight1 / totmass * 100)
		sodr2 := int(math.Ceil(wetWeight2 / totmass * 100))
		sodr3 := int(math.Ceil(wetWeight3 / totmass * 100))
		sodr4 := int(math.Ceil(wetWeight4 / totmass * 100))
		sodr5 := int(math.Ceil(wetWeight5 / totmass * 100))

		if sodr1+sodr2+sodr3+sodr4+sodr5 > 100 {
			if sodr5 > 0 {
				sodr5--
			} else if sodr4 > 0 {
				sodr4--
			} else if sodr3 > 0 {
				sodr3--
			} else if sodr2 > 0 {
				sodr2--
			} else {
				sodr1--
			}

		}

		sod1.Text = strconv.Itoa(sodr1)
		sod1.Refresh()
		sod2.Text = strconv.Itoa(sodr2)
		sod2.Refresh()
		sod3.Text = strconv.Itoa(sodr3)
		sod3.Refresh()
		sod4.Text = strconv.Itoa(sodr4)
		sod4.Refresh()
		sod5.Text = strconv.Itoa(sodr5)
		sod5.Refresh()

		totalwater.Text = strconv.Itoa(int(wta))
		totalwater.Refresh()
		totalmass.Text = strconv.Itoa(int(totmass))
		totalmass.Refresh()
	})
	calcByWeight.Move(fyne.NewPos(460, 250))
	calcByWeight.Resize(fyne.NewSize(55, 35))

	// Создаем контейнер для размещения элементов на экране
	content := container.NewWithoutLayout(
		zagol,
		zagol1, sod1, sod2, sod3, sod4, sod5,
		ves1, ves2, ves3, ves4, ves5,
		komp1, proc1, wat1,
		komp2, proc2, wat2,
		komp3, proc3, wat3,
		komp4, proc4, wat4,
		komp5, proc5, wat5,
		totalmassZag, totalmass,
		HumidityZag, Humidity,
		addWaterZag,
		calcByProc, calcByWeight, totalwater,
		version, telega,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
	tidyUp()

}
func tidyUp() {

}

func getHumid(x string) (string, string) {
	humidMap := map[string]string{
		"Берёзовый RUF":      "10",
		"Буковые пеллеты":    "8",
		"Пеллеты из лузги":   "15",
		"Отруби":             "25",
		"Пеллеты из люцерны": "10",
		"Пеллеты комбикорм":  "15",
		"Мел":                "0",
		"_________________":  "0",
	}
	soderMap := map[string]string{
		"Берёзовый RUF":      "30",
		"Буковые пеллеты":    "30",
		"Пеллеты из лузги":   "20",
		"Отруби":             "20",
		"Пеллеты из люцерны": "20",
		"Пеллеты комбикорм":  "20",
		"Мел":                "5",
		"_________________":  "0",
	}

	return humidMap[x], soderMap[x]

}

func calculation(fullweight, humidPercent, soder, procent int) (normalMassComponent, waterToAdd int) {

	totalWater := fullweight * humidPercent / 100 //общее количество воды
	compWeight := fullweight * soder / 100        // масса размоченного компонента
	waterInComponent := totalWater * soder / 100  //количество воды в данном компоненте
	dryMassComponent := (compWeight - waterInComponent)
	normalMassComponent = dryMassComponent * 100 / (100 - procent)
	waterToAdd = compWeight - normalMassComponent
	return normalMassComponent, waterToAdd

}
