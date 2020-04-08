#!/usr/bin/env php
<?php
/**
[case]
title=use ztf to run appium test
cid=0
pid=0

[group]
  1. check image element displayed attribute is >> true

[esac]
*/

require_once('vendor/autoload.php');

use Facebook\WebDriver\Remote\DesiredCapabilities;
use Facebook\WebDriver\Remote\RemoteWebDriver;
use Facebook\WebDriver\WebDriverBy;

class AndroidTest
{
    protected $webDriver;

    public function demo()
    {
        $capabilities = new DesiredCapabilities();
        $capabilities->setCapability("deviceName", "redmi");
        $capabilities->setCapability("platformName", "Android");

        // $capabilities->setCapability("app", "https://applitools.bintray.com/Examples/eyes-android-hello-world.apk");
		$capabilities->setCapability("app", '/Users/aaron/testing/res/eyes-android-hello-world.apk');

        $driver = RemoteWebDriver::create("http://172.16.13.1:4723/wd/hub", $capabilities);

            $driver->findElement(WebDriverBy::id("random_number_check_box"))->click();
            $driver->findElement(WebDriverBy::id("click_me_btn"))->click();

			$image = $driver->findElement(WebDriverBy::id("image"));
			print('>>' . $image->getAttribute('displayed') . "\n");

			$driver->quit();
    }
}

$test = new AndroidTest();
$test->demo();