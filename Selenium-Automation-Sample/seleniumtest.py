from selenium import webdriver
from  selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.action_chains import ActionChains
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC


class Driver():

    def __init__(self):
        self.driver = webdriver.Chrome()
        #url = "http://localhost:4444/wd/hub"
        #options = webdriver.ChromeOptions()
        #self.driver = webdriver.Remote(url, options.to_capabilities())


    def getpage(self,url):
        driver = self.driver
        driver.implicitly_wait(10)
        page = driver.get(url)
        driver.maximize_window()
        

    def selectregion(self,country):
        
        xpath = "//a[@title={0}]".format(country)
        driver = self.driver
        c = driver.find_element_by_xpath(xpath)
        c.click()

    def clickobj(self,xpath):
        driver = self.driver
        c = driver.find_element_by_xpath(xpath)
        c.click()
        
    
    def getdriver(self):
        return self.driver

    def newtab(self):
        self.driver.find_element_by_tag_name('body').send_keys(Keys.CONTROL + 't')
        self.driver.execute_script("window.open('');") 
    
    def getelements(self,xpath):
        elements = self.driver.find_elements_by_xpath(xpath)
        return elements

    def returnelement(self,xpath):
        element = self.driver.find_element_by_xpath(xpath)
        return element

    def movetoelement(self,xpath):
        action = ActionChains(self.driver)
        element = self.driver.find_element_by_xpath(xpath)
        action.move_to_element(element).click().perform()
        
    def expectedcondition(self,xpath):
        element = WebDriverWait(self.driver, 10).until(EC.presence_of_element_located((By.XPATH, xpath)))
        return element


    

def main():
    resultdict = {}
    a = Driver()
    b = a.getdriver()
    try:
        a.getpage("https://www.tesla.com/")
        a.selectregion("'United States'")
        a.clickobj('//label[@data-gtm-location="hamburger nav"]')
        xp = "//li[@class={0}]".format('"tds-menu-header-nav--list_item"')
        elements = a.getelements(str(xp))
        elements[2].click()
        a.newtab()

        b = a.getdriver()
        #print(b.window_handles)
        b.switch_to.window(b.window_handles[1])
        print(b.current_window_handle)
        a.getpage("https://www.gsmarena.com/")
        newel = a.returnelement('//input[@name="sSearch"]//parent::form//child::input')
        newel.click()
        newel.send_keys("lumia 950")
        newel.send_keys(Keys.RETURN)
        newel = a.returnelement("//div[@class='makers']//descendant::ul/li[2]")
        newel.click()
        iflem = a.expectedcondition("//div[contains(@class,'article')]//descendant::a[contains(text(),'Review')]")
        if iflem is None:
            print("the element was not found")
            resultdict["Test one"] = "Fail"
        else:
            resultdict["Test one"] = "Pass"

        a.movetoelement("//div[contains(@class,'article')]//descendant::a[contains(text(),'Review')]")
        b.get_screenshot_as_file("\\Users\\sudharsan.sivaram\\Downloads\\check.png")
        b.refresh()
        print(b.title)
        b.back()
       # b.close()
       # b.quit()
    except Exception as e:
        print(e)
       # b.close()
       # b.quit()

main()


