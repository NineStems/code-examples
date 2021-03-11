package calculate

import (
	models "go-rti-testing/src/models"
	"strconv"
)

/**Функция предназначена для вычисления равенства между список свойств и списком условий для предложения в продукте
*@param conditionProduct Условия цены
*@param conditionRule	 Условия из ограничений
 */
func conditionComparator(conditionProduct []models.RuleApplicability, conditionRule []models.Condition) (result bool) {
	var mapProductCondition = make(map[string]string)
	var countConditionProduct int

	for _, cProdCond := range conditionProduct {
		mapProductCondition[cProdCond.CodeName] = cProdCond.Value
	}
	for _, cRuleCond := range conditionRule {
		if val, hasEl := mapProductCondition[cRuleCond.RuleName]; hasEl {
			if val == cRuleCond.Value {
				countConditionProduct++
			}
		}
	}

	if len(conditionProduct) < len(conditionRule) {
		if countConditionProduct == len(conditionProduct) {
			result = true
		} else {
			result = false
		}
	} else if len(conditionProduct) > len(conditionRule) {
		if countConditionProduct == len(conditionRule) {
			result = true
		} else {
			result = false
		}
	} else {
		if countConditionProduct == len(conditionProduct) {
			result = true
		} else {
			result = false
		}
	}

	return
}

/**Функция предназначена для применения скидки к элементу, если проходит по условиям
*@param component 	Текущий компонент, в котором была цена, проходящая по условиям
*@param listPrices 	Цена, к которой требуется применить скидку
 */
func admitDiscount(component models.Component, listPrices *models.Price) {
	var discountValue float64
	for _, price := range component.Prices {
		if price.PriceType == models.PriceTypeDiscount {
			switch price.RuleApplicabilities[0].Operator {
			case models.OperatorGreaterThanOrEqual:
				for _, val := range listPrices.RuleApplicabilities {
					rulePrice, _ := strconv.Atoi(price.RuleApplicabilities[0].Value)
					listPrice, _ := strconv.Atoi(val.Value)
					if listPrice >= rulePrice {
						if discountValue < price.Cost {
							discountValue = price.Cost
						}
					}
				}
			case models.OperatorLessThanOrEqual:
				for _, val := range listPrices.RuleApplicabilities {
					rulePrice, _ := strconv.Atoi(price.RuleApplicabilities[0].Value)
					listPrice, _ := strconv.Atoi(val.Value)
					if listPrice <= rulePrice {
						if discountValue < price.Cost {
							discountValue = price.Cost
						}
					}
				}
			}

		}
	}
	if discountValue != 0 {
		listPrices.Cost = float64(listPrices.Cost * ((100 - discountValue) / 100))
	}
	return
}

/**Функция предназначена для обнуления полей, которые не нужны в JSON
*@param product Продукт для предложения, требующий обнуления лишних значений в его компонентах
 */
func clearExtraFieldsProduct(product *models.Product) {
	for indexComp, comp := range product.Components {
		for indexPrice := range comp.Prices {
			product.Components[indexComp].Prices[indexPrice].PriceType = ""
			product.Components[indexComp].Prices[indexPrice].RuleApplicabilities = nil
		}
	}
	return
}

func calculateOfferSum(components []models.Component) float64 {
	var returnSum float64
	for _, comp := range components {
		returnSum += comp.Prices[0].Cost
	}
	return returnSum
}

func Calculate(product *models.Product, conditions []models.Condition) (offer *models.Offer, err error) {
	var errResult error = nil
	var offerResult models.Offer
	var offerProduct models.Product
	var offerPrice models.Price
	var listComponent = make([]models.Component, 0)
	var listPrices = make([]models.Price, 0)
	var hasNeedContidion bool

	if product == nil || conditions == nil {
		return nil, errResult
	}

	components := product.Components

	for _, component := range components {
		hasNeedContidion = false
		listPrices = nil
		for _, price := range component.Prices {
			hasNeedContidion = conditionComparator(price.RuleApplicabilities, conditions)
			if hasNeedContidion && price.PriceType != models.PriceTypeDiscount {
				priceChange := price
				admitDiscount(component, &priceChange)
				listPrices = append(listPrices, priceChange)
				continue
			}
		}
		if listPrices != nil {
			listComponent = append(listComponent, models.Component{Name: component.Name, IsMain: component.IsMain, Prices: listPrices})
		}
	}
	offerProduct.Name = product.Name

	if len(listComponent) == 0 {
		listComponent = nil
	}

	for _, comp := range listComponent {
		if len(comp.Prices) > 1 {
			return nil, errResult
		}
	}

	offerProduct.Components = listComponent
	clearExtraFieldsProduct(&offerProduct)
	offerResult.Product = offerProduct
	offerPrice.Cost = calculateOfferSum(listComponent)
	offerResult.TotalCost = offerPrice

	return &offerResult, errResult
}
