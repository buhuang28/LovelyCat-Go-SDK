#include <stdint.h>
#include <windows.h>

#define CATEVENT extern int32_t __stdcall __declspec (dllexport)

extern int32_t NewAppInfo(int32_t);
//export LoadingInfo
CATEVENT LoadingInfo(int session){ return NewAppInfo(session);}

extern int32_t _EventInit();
extern int32_t _EventEnable();
extern int32_t _EventStop();
extern int32_t _EventLogin(char *,char *,int32_t);
extern int32_t _EventGroupMsg(char *,int32_t,char *,char *,char *,char *,char *,char *);
extern int32_t _EventFriendMsg(char *,int32_t, char *, char *, char *, char *);
extern int32_t _EventReceivedTransfer(char *,char *, char *, char *,char *,char *);
extern int32_t _EventScanCashMoney(char *, char *, char *,char *,char *);
extern int32_t _EventFriendVerify(char *, char *, char *, char *,char *);
extern int32_t _EventGroupMemberAdd(char *, char *, char *,char *);
extern int32_t _EventGroupMemberDecrease(char *, char *, char *, char *);
extern int32_t _EventSysMsg(char *, int32_t, char *);
extern int32_t _Menu();

//export EventInit
CATEVENT EventInit(){ return _EventInit();}
//export EventEnable
CATEVENT EventEnable(){ return _EventEnable();}
//export EventStop
CATEVENT EventStop(){ return _EventStop();}
//export EventLogin
CATEVENT EventLogin(char *rob_wxid,char *rob_wxname,int32_t type){ return _EventLogin(rob_wxid,rob_wxname,type);}
//export EventGroupMsg
CATEVENT EventGroupMsg(char *robot_wxid,int32_t type,char *from_wxid,char *from_name,char *final_from_wxid,char *final_from_name,char *to_wxid,char *Msg){ return _EventGroupMsg(robot_wxid,type,from_wxid,from_name,final_from_wxid,final_from_name,to_wxid,Msg);}
//export EventFriendMsg
CATEVENT EventFriendMsg(char *robot_wxid,int32_t type, char *from_wxid, char *from_name, char *to_wxid, char *Msg){ return _EventFriendMsg(robot_wxid,type,from_wxid,from_name,to_wxid,Msg);}
//export EventReceivedTransfer
CATEVENT EventReceivedTransfer(char *robot_wxid,char *from_wxid, char *from_name, char *to_wxid,char *money,char *Msg_json){ return _EventReceivedTransfer(robot_wxid,from_wxid,from_name,to_wxid,money,Msg_json);}
//export EventScanCashMoney
CATEVENT EventScanCashMoney(char *robot_wxid, char *pay_wxid, char *pay_name,char *money,char *Msg_json){ return _EventScanCashMoney(robot_wxid,pay_wxid,pay_name,money,Msg_json);}
//export EventFriendVerify
CATEVENT EventFriendVerify(char *robot_wxid, char *from_wxid, char *from_name, char *to_wxid,char *Msg_json){ return _EventFriendVerify(robot_wxid,from_wxid,from_name,to_wxid,Msg_json);}
//export EventGroupMemberAdd
CATEVENT EventGroupMemberAdd(char *robot_wxid, char *from_wxid, char *from_name,char *Msg_json){ return _EventGroupMemberAdd(robot_wxid,from_wxid,from_name,Msg_json);}
//export EventGroupMemberDecrease
CATEVENT EventGroupMemberDecrease(char *robot_wxid, char *from_wxid, char *from_name, char *Msg_json){ return _EventGroupMemberDecrease(robot_wxid,from_wxid,from_name,Msg_json);}
//export EventSysMsg
CATEVENT EventSysMsg(char *robot_wxid, int32_t type, char *Msg_json){ return _EventSysMsg(robot_wxid,type,Msg_json);}
//export Menu
CATEVENT Menu(){ return _Menu();}
