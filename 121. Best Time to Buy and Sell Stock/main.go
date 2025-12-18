func maxProfit(prices []int) int {
    var minPrice, maxProfit int = prices[0], 0
    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        }
        if profit := price - minPrice; profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}