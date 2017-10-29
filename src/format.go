package main 

import (
	"encoding/gob"
	"fmt"
	"os"
	"runtime"
)

type Trail struct{
	pondsTraveled[12] int
	pondsGrade[12] int
	remark string 
}

type Fish struct {
	fishName string
	description string
	travelPath Trail
	incomingPond int 
	outgoingPond int 
}

type DirtyFish struct {
	fishName string
	dirtyCode string
}

type Pond struct {
	pondNumber int
	description string
	fish[13] Fish
}

type FishofPond struct {
	pond[13] Pond
	dirtyLane[13] DirtyFish

}

func initializePool() FishofPond{
	pool := FishofPond{}
	return pool
}

func initializeAllPonds(pool *FishofPond) {
	pool.pond[0].pondNumber = 0
	pool.pond[0].description = "IncomingFilter"
	pool.pond[1].pondNumber = 1
	pool.pond[1].description = "ManagementStudy"
	pool.pond[2].pondNumber = 2
	pool.pond[2].description = "BusinessModel"
	pool.pond[3].pondNumber = 3
	pool.pond[3].description = "ValuationStudy"
	pool.pond[4].pondNumber = 4
	pool.pond[4].description = "FirstDecisionPool"
	pool.pond[5].pondNumber = 5
	pool.pond[5].description = "LongTermInvestment"
	pool.pond[6].pondNumber = 6
	pool.pond[6].description = "ShortTermInvestment"
	pool.pond[7].pondNumber = 7
	pool.pond[7].description = "TakeLossBookProfit"
	pool.pond[8].pondNumber = 8
	pool.pond[8].description = "SecondDecisionPool"
	pool.pond[9].pondNumber = 9
	pool.pond[9].description = "ReInvestLongTerm"
	pool.pond[10].pondNumber = 10
	pool.pond[10].description = "ReTradeShortTerm"
	pool.pond[11].pondNumber = 11
	pool.pond[11].description = "ProfitList"
    pool.pond[12].pondNumber = 12
	pool.pond[12].description = "LossList"
    
}

func addFishInPondZero(pond *Pond,name string,desc string) {
	if pond.pondNumber == 0 {
		var i=0
	    for pond.fish[i].fishName != "" {
	    	i = i + 1
	    	fmt.Println("Finding place for fish at ",i)
	    }
        var spaceInPond = len(pond.fish) - i

        if spaceInPond < 1 {
        	fmt.Println("No space left in pond")
        	os.Exit(0)
        }
	    fmt.Println("Space left in pond ",spaceInPond)
	    fmt.Println("Adding fish",i)

	    pond.fish[i].fishName = name 
	    pond.fish[i].description = desc
	    pond.fish[i].outgoingPond = 1
    
	}
}


func main() {
	//Initialize complete pool with 12 ponds and 12 fishes in each
	var pool = initializePool()
	initializeAllPonds(&pool)
	//pool := addFishInPondZero(pool.pond[0],)
	addFishInPondZero(&pool.pond[0],"DHFL","HousingFinance")
	//fmt.Println(pool.pond[0])

	addFishInPondZero(&pool.pond[0],"MGL","CleanEnergy")
	addFishInPondZero(&pool.pond[0],"Suzlon","CleanEnergy")
	addFishInPondZero(&pool.pond[0],"Fiem","AutoAncillary")
	addFishInPondZero(&pool.pond[0],"Mindtree","InformationTechnology")
	addFishInPondZero(&pool.pond[0],"Selan Oil Exploration","Energy")
	fmt.Println(pool.pond[0])
	/*fmt.Println(pool.pond[2].fish[1])
	pool.pond[2].pondNumber = 2
	pool.pond[2].description = "BusinessModel"
    pool.pond[2].fish[1].fishName = "DHFL"
    pool.pond[2].fish[1].description = "Housing finance"
    pool.pond[2].fish[1].travelPath.pondsTraveled[1] = 1
    pool.pond[2].fish[1].travelPath.pondsGrade[1] = 1
    pool.pond[2].fish[1].travelPath.remark = "Good management"
    pool.pond[2].fish[1].incomingPond = 1
    pool.pond[2].fish[1].outgoingPond = 3
    */
    //fmt.Println(pool.pond[2].fish[1])
    //pool.pond[1]
    /*
	var fishName = "DHFL"
	var incomingPond = 0
	var outgoingPond = 2
	var payload = 0
	var description = "HF"

	fish := Fish{fishName,incomingPond,outgoingPond,payload,description}
	pool := FishofPond{}
	fmt.Println(pool)*/
	
}

// Encode via Gob to file
func Save(path string, object interface{}) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
 }

// Decode Gob file
func Load(path string, object interface{}) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func Check(e error) {
	if e != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(line, "\t", file, "\n", e)
		os.Exit(1)
	}
}
