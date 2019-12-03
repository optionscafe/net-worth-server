//
// Date: 12/2/2019
// Author(s): Spicer Matthews (spicer@options.cafe)
// Copyright: 2019 Cloudmanic Labs, LLC. All rights reserved.
//

package etrade

import (
	"testing"

	"github.com/nbio/st"
)

//
// ParseDigestEmail will parse the HTML of your daily etrade digest email.
// We do this because etrade's API sucks and makes users login every day.
//
func TestParseDigestEmail01(t *testing.T) {
	// Sample HTML from the etrade digest email.
	html := getTestHTML()

	parsed, err := ParseDigestEmail(html)

	// Check expected results.
	st.Expect(t, err, nil)
	st.Expect(t, parsed.Balance, 394004.72)
}

//
// getTestHTML get test html
//
func getTestHTML() string {
	return `
		<html xmlns:fo="http://www.w3.org/1999/XSL/Format">
		<head>
		<META http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
		<meta name="format-detection" content="telephone=no">
		<meta name="x-apple-disable-message-reformatting">
		</head>
		<body style="padding:0;">
		<div class="smart-alert">
		<link href="https://cdn.etrade.net/1/1d/aempros/etc/designs/smartalerts/styles/smartalerts.css" rel="stylesheet" type="text/css">
		<table align="center" width="100%" cellpadding="0" cellspacing="0" border="0" bgcolor="#ffffff" style="background-color:#ffffff;">
		<tbody>
		<div class="previewHeader" style="display:none;font-size:0;line-height:0;max-height:0;mso-hide:all">Portfolio Digest As of
		    Mon Dec  2 21:10:59 2019Portfolio: XXXX--9323  Personal Trading
		        &nbsp;
		      Market Value:
		       $394,004.72 (
		       -$23.67/-1.91%)
		        &nbsp;
		      SymbolCurrentPr</div>
		<tr>
		<td valign="top" align="center" bgcolor="#ffffff">
		<table width="640" border="0" cellspacing="0" cellpadding="0" style="table-layout: fixed;" class="em_main_table">
		<tbody>
		<tr>
		<td width="20px" bgcolor="#ffffff">&nbsp;</td><td valign="top" align="center">
		<table width="600" border="0" cellspacing="0" cellpadding="0" class="em_main_table" bgcolor="#ffffff">
		<tbody>
		<tr>
		<td height="12" style="line-height:1px; font-size:1px;">&nbsp;</td>
		</tr>
		<tr>
		<td valign="top" align="center">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center" class="em_main_table">
		<tbody>
		<tr>
		<td valign="top" align="left"><a href="https://us.etrade.com/home" target="_blank" title="E*TRADE" alias="preheader_logo" style="text-decoration:none; color:#51ae31;"><img src="https://cdn.etrade.net/1/1d/aempros/content/dam/etrade/smart-alerts/en_US/images/etrade_logo_small.png" width="241" height="auto" alt="E*TRADE" style="display:block; font-family:Arial, sans-serif; font-size:25px; line-height:49px; color:#51ae31; font-weight:bold; text-align:center;height:auto;width:241px" border="0" class="em_img"></a></td><td width="100" class="em_side">&nbsp;</td><td valign="bottom" align="right">
		<table width="315" border="0" cellspacing="0" cellpadding="0" align="right" class="em_hide">
		<tbody>
		<tr>
		<td valign="bottom" align="right" style="font-family: Arial, sans-serif; font-size:13px;  line-height:18px; color:#555555; text-align:right;font-weight:normal;" class="em_hide"><span class="em_hide"><a href="https://us.etrade.com/e/t/alerts/Alertinbox" target="_blank" alias="preheader" style="text-decoration:none; color:#6633cc;"><span style="color:#6633cc; text-decoration:none;">Alert Inbox</span></a>&nbsp;&nbsp;|&nbsp;&nbsp;</span><a href="https://us.etrade.com/e/t/user/login" target="_blank" mcdisabled-title="" alias="preheader" style="text-decoration:none; color:#6633cc;"><span style="color:#6633cc; text-decoration:none;">Log on</span></a></td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		<tr>
		<td style="line-height:0px; font-size:0px;" height="0">&nbsp;</td>
		</tr>
		<tr>
		<td height="20" style="line-height:1px; font-size:1px;">&nbsp;</td>
		</tr>
		</tbody>
		</table>
		</td><td width="20px" bgcolor="#ffffff">&nbsp;</td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		<!--[if !mso]>-->
		<tr>
		<td valign="top" align="center" bgcolor="#ffffff">
		<div class="hide_disktop" style="display:none;width:0;overflow:hidden;">
		<table width="620" border="0" cellspacing="0" cellpadding="0" style="display:none;table-layout: fixed; background-color:#ffffff;" class="em_main_table_hero">
		<tbody>
		<tr>
		<td width="20px">&nbsp;</td><td height="0" valign="top" align="center" width="100%">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
		<tbody>
		<tr>
		<td valign="top" align="left" width="45%" style="font-family: Arial, sans-serif; font-size:13px; line-height:18px; color:#555; text-align:left;font-weight:normal;"><a href="https://us.etrade.com/e/t/alerts/Alertinbox" target="_blank" alias="preheader" style="text-decoration:none; color:#6633cc;"><span class="visited" style="color:#6633cc; text-decoration:none;">Alert Inbox</span></a></td><td valign="top" width="45%" align="right" style="font-family: Arial, sans-serif; font-size:13px; line-height:18px; color:#555; text-align:right;font-weight:normal;"><a href="https://us.etrade.com/e/t/user/login" target="_blank" mcdisabled-title="" alias="preheader" style="text-decoration:none; color:#555;"><span class="visited" style="color:#6633cc; text-decoration:none;">Log On</span></a></td>
		</tr>
		</tbody>
		</table>
		</td><td width="20px">&nbsp;</td>
		</tr>
		<tr>
		<td style="line-height:0px; font-size:0px;" height="20px">&nbsp;</td>
		</tr>
		</tbody>
		</table>
		</div>
		</td>
		</tr>
		<!--<![endif]-->
		<tr>
		<td valign="top" align="center" height="15px" bgcolor="#19223c">
		    	&nbsp;
		    </td>
		</tr>
		<tr>
		<td align="center" bgcolor="#ffffff">
		<table valign="top" cellpadding="0" cellspacing="0" border="0" width="640" class="em_wrapper" bgcolor="#ffffff">
		<tr>
		<td height="30px">&nbsp;</td>
		</tr>
		</table>
		<table valign="top" align="center" cellpadding="0" cellspacing="0" border="0" width="640px" class="em_wrapper" bgcolor="#ffffff">
		<tr>
		<td width="20px">
		                    &nbsp;
		                </td><td valign="top" align="left" width="600" class="body-content" style="font-family: Arial, sans-serif;font-size: 18px;line-height: 24px;padding-bottom: 15px;color: #555555;">
		<table width="600" class="em_wrapper">
		<tr>
		<td align="center"><span style="font-size: 20px;"><b>Portfolio Digest</b></span>
		<br>
		<span class="coF1"> As of
		    Mon Dec  2 21:10:59 2019</span></td>
		</tr>
		</table>
		<table border="0" cellspacing="0" cellpadding="0" width="600">
		<tr>
		<td height="45"><strong>Portfolio: </strong>XXXX--9323  Personal Trading</td><td align="right" colspan="2"></td>
		</tr>
		</table>
		<table border="0" cellspacing="0" cellpadding="0" width="600" style="border-collapse: initial;">
		<tr>
		<td style="border: 1px solid #ccc; border-bottom: 0;">
		<table border="0" cellpadding="0" cellspacing="0" width="100%">
		<tr>
		<td height="29" nowrap="nowrap">
		        &nbsp;
		      <b>Market Value:
		       $394,004.72 (
		       <font color="#ff0000">-$23.67</font></b>/<b><font color="#ff0000">-1.91%</font>) </b></td><td align="right" nowrap="nowrap">
		        &nbsp;
		      </td>
		</tr>
		</table>
		</td>
		</tr>
		<tr>
		<td>
		<table border="0" cellpadding="3" cellspacing="1" width="100%" bgcolor="#cccccc" style="border-collapse: initial;">
		<tr bgcolor="#E3E3E3">
		<td align="center" colspan="2"><b>Symbol</b></td><td align="center"><b>Current<br>Price</b></td><td colspan="2">
		<table width="100%">
		<tr align="center">
		<td colspan="2"><b>Today's Change</b></td>
		</tr>
		<tr align="center">
		<td width="50%"><b>$</b></td><td width="50%"><b>%</b></td>
		</tr>
		</table>
		</td><td align="center"><b>Today's<br>Gain</b></td><td align="center"><b>Qty</b></td><td align="center"><b>Price Paid</b></td><td colspan="2">
		<table width="100%">
		<tr align="center">
		<td colspan="2"><b>Total Gain</b></td>
		</tr>
		<tr align="center">
		<td width="50%"><b>$</b></td><td width="50%"><b>%</b></td>
		</tr>
		</table>
		</td>
		</tr>
		<tr bgcolor="#ffffff" align="right">
		<td><b> <a href="https://us.etrade.com/e/t/invest/quotesresearch?traxui=P_MGR&amp;sym=TEAM" target="_blank">TEAM</a>
		        </b></td><td align="center"> <a href="https://us.etrade.com/e/t/invest/socreateentry?ordertype=1&amp;symbol=TEAM" target="_blank">
		Trade
		</a></td><td><b>120.84
		        </b></td><td><font color="#ff0000">-$6.27</font>
		       </td><td><font color="#ff0000">-4.93%</font></td><td><font color="#ff0000">-$10.15</font>
		       </td><td>4
		       </td><td>$123.67
		       </td><td><font color="#ff0000">-$11.33</font>
		       </td><td><font color="#ff0000">-2.29%</font>
		       </td>
		</tr>
		<tr bgcolor="#ffffff" align="right">
		<td><b> <a href="https://us.etrade.com/e/t/invest/quotesresearch?traxui=P_MGR&amp;sym=CYBR" target="_blank">CYBR</a>
		        </b></td><td align="center"> <a href="https://us.etrade.com/e/t/invest/socreateentry?ordertype=1&amp;symbol=CYBR" target="_blank">
		Trade
		</a></td><td><b>120.10
		        </b></td><td><font color="#ff0000">-$2.45</font>
		       </td><td><font color="#ff0000">-2.00%</font></td><td><font color="#ff0000">-$4.90</font>
		       </td><td>2
		       </td><td>$124.22
		       </td><td><font color="#ff0000">-$8.24</font>
		       </td><td><font color="#ff0000">-3.32%</font>
		       </td>
		</tr>
		<tr bgcolor="#ffffff" align="right">
		<td><b> <a href="https://us.etrade.com/e/t/invest/quotesresearch?traxui=P_MGR&amp;sym=CPRT" target="_blank">CPRT</a>
		        </b></td><td align="center"> <a href="https://us.etrade.com/e/t/invest/socreateentry?ordertype=1&amp;symbol=CPRT" target="_blank">
		Trade
		</a></td><td><b>87.96
		        </b></td><td><font color="#ff0000">-$1.04</font>
		       </td><td><font color="#ff0000">-1.17%</font></td><td><font color="#ff0000">-$3.12</font>
		       </td><td>3
		       </td><td>$90.19
		       </td><td><font color="#ff0000">-$6.70</font>
		       </td><td><font color="#ff0000">-2.48%</font>
		       </td>
		</tr>
		<tr bgcolor="#ffffff" align="right">
		<td><b> <a href="https://us.etrade.com/e/t/invest/quotesresearch?traxui=P_MGR&amp;sym=WORK" target="_blank">WORK</a>
		        </b></td><td align="center"> <a href="https://us.etrade.com/e/t/invest/socreateentry?ordertype=1&amp;symbol=WORK" target="_blank">
		Trade
		</a></td><td><b>22.51
		        </b></td><td><font color="#ff0000">-$0.31</font>
		       </td><td><font color="#ff0000">-1.36%</font></td><td><font color="#ff0000">-$3.10</font>
		       </td><td>10
		       </td><td>$22.25
		       </td><td><font color="#669900">$2.60</font>
		       </td><td><font color="#669900">1.17%</font>
		       </td>
		</tr>
		<tr bgcolor="#ffffff">
		<td align="right" colspan="5" bgcolor="#cccccc"><b>Totals</b></td><td align="right"><font color="#ff0000">-$21.27</font></td><td colspan="2">
		      </td><td align="right"><font color="#ff0000">-$23.67</font></td><td align="right"><font color="#ff0000">-1.91%</font></td>
		</tr>
		</table>
		</td>
		</tr>
		</table>
		<br>
		<br>
		<table width="600" border="0" cellspacing="0" cellpadding="0" align="center" class="em_wrapper" bgcolor="#e5e5e5">
		<tr>
		<td class="em_wrapper" bgcolor="#e5e5e5" valign="top" align="left" style="font-family:Arial,sans-serif; font-size:18px; line-height:24px; color:#000000; background-color:#e5e5e5; padding:5px 10px 5px 10px; text-align:left; font-weight:bold;" width="100%"> Portfolio news </td>
		</tr>
		</table>
		<table width="600" class="em_wrapper">
		<tr>
		<td>
		<table border="0" cellspacing="0" cellpadding="4" width="100%">
		<tr bgcolor="#f4f4f4">
		<td width="23" align="left">Date</td><td align="left">News Stories</td><td align="left">Source</td>
		</tr>
		<tr>
		<td colspan="3"><b><a href="https://us.etrade.com/e/t/invest/quotesresearch?sym=TEAM">TEAM</a></b></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		23</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;09:34 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191123MW000089&amp;provider=MarketWatch.com">MW UPDATE: U.S. stocks look fully priced -- where to put your money now may surprise you</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr>
		<td colspan="2" align="right"><a href="https://us.etrade.com/e/t/invest/SingleSymbolHeadline?symbol=TEAM&amp;latest=11/23/2019-09:34">more headlines for TEAM...</a></td>
		</tr>
		<tr>
		<td colspan="3">
		<table border="0" cellpadding="0" cellspacing="0" width="100%">
		<tr bgcolor="#000000">
		<td><img src="https://cdn.etrade.net/1/7d/images/spacer.gif" width="500" height="1"></td>
		</tr>
		</table>
		</td>
		</tr>
		<tr>
		<td colspan="3"><b><a href="https://us.etrade.com/e/t/invest/quotesresearch?sym=CYBR">CYBR</a></b></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		25</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;07:00 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191125DN002369&amp;provider=DowJones">*DJ CyberArk Software Names Matthew Cohen as Revenue Chief</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;07:00 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191125SN004357&amp;provider=BusinessWire">CyberArk Names Matthew Cohen Chief Revenue Officer</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Business Wire</nobr></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		14</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;07:05 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191114DN005129&amp;provider=DowJones">*DJ CyberArk Software Ltd. Announces Pricing of Private Offering of $500M of 0% Convertible Senior Notes Due 2024</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr>
		<td colspan="2" align="right"><a href="https://us.etrade.com/e/t/invest/SingleSymbolHeadline?symbol=CYBR&amp;latest=11/14/2019-07:05">more headlines for CYBR...</a></td>
		</tr>
		<tr>
		<td colspan="3">
		<table border="0" cellpadding="0" cellspacing="0" width="100%">
		<tr bgcolor="#000000">
		<td><img src="https://cdn.etrade.net/1/7d/images/spacer.gif" width="500" height="1"></td>
		</tr>
		</table>
		</td>
		</tr>
		<tr>
		<td colspan="3"><b><a href="https://us.etrade.com/e/t/invest/quotesresearch?sym=CPRT">CPRT</a></b></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		22</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;09:01 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191122DN004285&amp;provider=DowJones">DJ Copart Price Target Raised to $100.00/Share From $92.00 by SunTrust Robinson Humphrey</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;07:47 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191122MW000165&amp;provider=MarketWatch.com">MW Copart stock price target raised to $100 from $92 at SunTrust RH</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		21</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;
		04:24 PM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191121DN009915&amp;provider=DowJones">DJ Copart's Management on Q1 2020 Results -- Earnings Call Transcript &gt;CPRT</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr>
		<td colspan="2" align="right"><a href="https://us.etrade.com/e/t/invest/SingleSymbolHeadline?symbol=CPRT&amp;latest=11/21/2019-16:24">more headlines for CPRT...</a></td>
		</tr>
		<tr>
		<td colspan="3">
		<table border="0" cellpadding="0" cellspacing="0" width="100%">
		<tr bgcolor="#000000">
		<td><img src="https://cdn.etrade.net/1/7d/images/spacer.gif" width="500" height="1"></td>
		</tr>
		</table>
		</td>
		</tr>
		<tr>
		<td colspan="3"><b><a href="https://us.etrade.com/e/t/invest/quotesresearch?sym=WORK">WORK</a></b></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		25</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;
		01:15 PM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191125DN008404&amp;provider=DowJones">DJ Burned by WeWork, SoftBank CEO Masayoshi Son Focuses on Cash Flow -- Barrons.com</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		20</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;08:03 AM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191120MW000116&amp;provider=MarketWatch.com">MW UPDATE: Microsoft Teams is making Slack investors nervous with its growth</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<TR>
		<td align="left" colspan="2" style="padding: 10px 0px 0px 0px;"><strong>
		Nov
		19</strong></td>
		</TR>
		<tr valign="top">
		<td align="center" nowrap="nowrap">&nbsp;&nbsp;&nbsp;
		05:32 PM
		&nbsp;&nbsp;&nbsp;</td><td><a href="https://us.etrade.com/e/t/invest/Story?ID=STORYID%3D20191119MW000510&amp;provider=MarketWatch.com">MW UPDATE: Slack shares slump as Microsoft continues to make corporate inroads</a>&nbsp;&nbsp;</td><td nowrap="nowrap"><nobr>Reuters World News</nobr></td>
		</tr>
		<tr>
		<td colspan="2" align="right"><a href="https://us.etrade.com/e/t/invest/SingleSymbolHeadline?symbol=WORK&amp;latest=11/19/2019-17:32">more headlines for WORK...</a></td>
		</tr>
		</table>
		</td>
		</tr>
		</table>
		<style> @media only screen and (max-width: 700px) { table[class=em_wrapper_mobile] { display:block !important; width:100% !important; } table[class=em_wrapper_desktop] { display:none !important; width:100% !important; } } .em_wrapper_mobile, .em_wrapper_desktop{ font-size: 16px !important; color: #000000;} .first_td { padding: 0 0 2px 15px; } .last_td { padding-bottom: 15px;} .smart-alert * {line-height: 24px;} .mob_date{ font-size: 14px; font-weight: bold;} .body-content , .sa-inbox-body{font-size: 12px !important; color: #000000 !important; } </style>
		</td><td width="20px">
		                    &nbsp;
		                </td>
		</tr>
		</table>
		</td>
		</tr>
		<tr>
		<td valign="top" align="center" bgcolor="#eeeeee"><a name="disclaimer" id="disclaimer"></a></td>
		</tr>
		<tr>
		<td valign="top" align="center" style="border-top:5px solid #dbdbdb;" bgcolor="#eeeeee"></td>
		</tr>
		<tr>
		<td bgcolor="#eeeeee" align="center" valign="top">
		<table width="640" border="0" cellspacing="0" cellpadding="0" class="em_main_table_disc">
		<tbody>
		<tr>
		<td align="center" valign="top">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
		<tbody>
		<tr>
		<td align="center" valign="top">
		<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
		<tbody>
		<tr>
		<td align="left" valign="top" class="disclaimer_text" style="font-family: Arial, sans-serif; font-size: 11px; line-height: 16px; padding: 20px; color: #555555;">
		<p>Change or manage your alert <a href="https://us.etrade.com/e/t/alerts/setdelivery" target="_blank" style="color: #555555;text-decoration: underline; ">delivery preferences</a>.</p>
		<p>(c) 2019 E*TRADE Securities LLC, Member <a href="http://www.finra.org/" target="_blank" style="color: #555555;text-decoration: underline; ">FINRA</a>/<a href="http://www.sipc.org/" target="_blank" style="color: #555555;text-decoration: underline; ">SIPC</a>. All rights reserved. <a href="https://us.etrade.com/l/f/s/brokerage" target="_blank" style="color: #555555;text-decoration: underline; ">Important disclosures</a>.</p>
		</td>
		</tr>
		<tr></tr>
		<tr>
		<td height="28" class="em_height" style="line-height:1px;font-size:1px;">&nbsp;
		                                                        <div class="em_non" style=" white-space:nowrap; font:20px courier; color:#eeeeee;">
		                                                            &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;</div>
		</td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		</tbody>
		</table>
		</td>
		</tr>
		</tbody>
		</table>
		</div>
		</body>
		</html>
`
}

/* End File */
