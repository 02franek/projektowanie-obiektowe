import { Builder, By, until } from "selenium-webdriver";
import assert from "node:assert";

describe("Rejestracja Użytkownika (Selenium)", function () {
  this.timeout(30000);
  let driver;

  before(async function () {
    driver = await new Builder().forBrowser("chrome").build();
  });

  after(async function () {
    if (driver) {
      await driver.quit();
    }
  });

  beforeEach(async function () {
    await driver.get("http://localhost:5173/register");

    await driver.wait(
      until.elementLocated(By.xpath("//h2[contains(text(), 'Rejestracja')]")),
      5000,
    );
  });

  it("1. Powinien zablokować rejestrację i wyświetlić błędy dla pustych pól obowiązkowych", async function () {
    let submitBtn = await driver.findElement(By.css('button[type="submit"]'));
    await submitBtn.click();

    let usernameError = await driver.wait(
      until.elementLocated(
        By.xpath(
          "//span[contains(text(), 'Nazwa użytkownika jest wymagana.')]",
        ),
      ),
      5000,
    );
    let emailError = await driver.findElement(
      By.xpath("//span[contains(text(), 'Adres e-mail jest wymagany.')]"),
    );
    let passwordError = await driver.findElement(
      By.xpath("//span[contains(text(), 'Hasło jest wymagane.')]"),
    );

    assert.strictEqual(await usernameError.isDisplayed(), true);
    assert.strictEqual(await emailError.isDisplayed(), true);
    assert.strictEqual(await passwordError.isDisplayed(), true);
  });

  it("2. Powinien zgłosić błąd przy niepoprawnym formacie adresu e-mail", async function () {
    await driver.findElement(By.name("username")).sendKeys("TestowyUser");
    await driver
      .findElement(By.name("email"))
      .sendKeys("niepoprawny-adres-bez-malpy");
    await driver.findElement(By.name("password")).sendKeys("haslo123");

    let submitBtn = await driver.findElement(By.css('button[type="submit"]'));
    await submitBtn.click();

    let emailFormatError = await driver.wait(
      until.elementLocated(
        By.xpath(
          "//span[contains(text(), 'Wprowadzono niepoprawny format adresu e-mail.')]",
        ),
      ),
      5000,
    );

    assert.strictEqual(await emailFormatError.isDisplayed(), true);
  });

  it("3. Powinien przejść pomyślnie przy poprawnych danych", async function () {
    await driver.findElement(By.name("username")).sendKeys("PoprawnyUser");
    await driver.findElement(By.name("email")).sendKeys("test@domena.com");
    await driver.findElement(By.name("password")).sendKeys("Haslo123321!.");

    let submitBtn = await driver.findElement(By.css('button[type="submit"]'));
    await submitBtn.click();

    let successMessage = await driver.wait(
      until.elementLocated(
        By.xpath(
          "//div[contains(text(), 'Rejestracja przebiegła pomyślnie!')]",
        ),
      ),
      5000,
    );

    assert.strictEqual(await successMessage.isDisplayed(), true);
  });
});
