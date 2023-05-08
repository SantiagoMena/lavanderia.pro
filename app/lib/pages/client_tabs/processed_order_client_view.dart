import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:lavanderiapro/auth/register_business.dart';
import 'package:lavanderiapro/services/get_order_service.dart';
import 'package:lavanderiapro/services/get_profile_service.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ProcessedOrderClient extends StatefulWidget {
   const ProcessedOrderClient({super.key, this.orderId});

   final String? orderId;

  @override
  State<ProcessedOrderClient> createState() => _ProcessedOrderClientState();
}

class _ProcessedOrderClientState extends State<ProcessedOrderClient> {

  @override
  Widget build(BuildContext context) {

    return Scaffold(
        appBar: AppBar(
        title: Text("Order to Business Name"),
      ),
      body: LayoutBuilder(
      builder: (BuildContext context, BoxConstraints viewportConstraints) {
      return Align(
        alignment: Alignment.topCenter,
        child: SingleChildScrollView(
              reverse: true,
              child: FutureBuilder(
                future: SharedPreferences.getInstance(),
                builder: (contextSharedPreference, snapshot) {
                  if(snapshot.hasData) {
                    String token = snapshot.data!.getString('token') ?? "";
                    String orderId = widget.orderId ?? "";

                    return FutureBuilder(
                    future: getOrder(token, orderId),
                    builder: (contextOrder, snapshotOrder) {
                      if(snapshotOrder.hasData){
                        return Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        mainAxisAlignment: MainAxisAlignment.start,
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          Padding(
                            padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                            child: Text("🧾", style: TextStyle(fontSize: 70),)
                          ),
                          Padding(
                              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 16),
                              child: Text(snapshotOrder.data!.address?.address ?? "")
                          ),
                          Padding(padding: EdgeInsets.only(bottom: MediaQuery.of(context).viewInsets.bottom))
                        ],
                       );
                      }
                      else {
                        return const Text("Order Not Found");
                      }
                    }
                  );
                  }
                  else {
                    return const CircularProgressIndicator();
                  }
                }
              ),
              ),
        );
      }
      ),
    );
  }
}
