func maximumWealth(accounts [][]int) int {
    
    wealth := 0
    
    for _,  val := range accounts {
        
//         calculating wealth for each acc
        maxWealth := 0
        for _, v := range val {
            maxWealth += v
        }
        
//         update maxwealth if any have greater wealth
        if maxWealth > wealth {
            wealth = maxWealth
        }
    }
    
    return wealth
    
}