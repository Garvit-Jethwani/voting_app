WebDriver driver = new ChromeDriver();
String url = System.getenv("ROOST_SVC_URL");
driver.get(url);

Assert.assertTrue(driver.getPageSource().contains("Election Commission Admin Portal"));

WebElement addCandidateButton = driver.findElement(By.id("add-candidate-button"));
addCandidateButton.click();

Assert.assertTrue(driver.getPageSource().contains("Add your candidates for election of K8s distribution here"));
Assert.assertTrue(driver.getPageSource().contains("Candidate's Name"));
Assert.assertTrue(driver.getPageSource().contains("Candidate's Picture"));